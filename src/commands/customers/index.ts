import DodoPayments from 'dodopayments';
import { isDodoPaymentsAPIError } from '../../utils/error';
import { paginationTip } from '../../utils/tips';
import type { CommandContext } from '../../tui/CommandContext';

export async function handleCustomers(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading customers…' });
    try {
      const customers = (
        await client.customers.list({
          page_number: pageNum - 1,
          page_size: 100,
        })
      ).items;

      ctx.removeBlock(spinnerId);
      if (customers.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }
      ctx.addBlock({
        type: 'table',
        data: customers.map((c) => ({
          customer_id: c.customer_id,
          name: c.name,
          email: c.email,
          phone_number: c.phone_number,
        })),
      });
      ctx.addBlock({ type: 'streaming', text: paginationTip('/customers list') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load customers. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'create') {
    const name = await ctx.promptInput('Name');
    const email = await ctx.promptInput('Email');
    const phone = await ctx.promptInput('Phone number');

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Creating customer…' });
    try {
      const creation = await client.customers.create({
        name,
        email,
        phone_number: phone.trim() !== '' ? phone : null,
      });

      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Customer created.' });
      ctx.addBlock({
        type: 'detail',
        data: {
          customer_id: creation.customer_id,
          name: creation.name,
          email: creation.email,
          phone_number: creation.phone_number,
        },
      });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't create customer. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'update') {
    const customer_id = args[0];
    if (!customer_id) {
      ctx.addBlock({ type: 'error', message: 'Customer ID required. Usage: /customers update <id>' });
      return;
    }

    const fetchSpinner = ctx.addBlock({ type: 'spinner', label: 'Loading customer…' });
    let existingInfo;
    try {
      existingInfo = await client.customers.retrieve(customer_id);
      ctx.removeBlock(fetchSpinner);
    } catch (e: any) {
      ctx.removeBlock(fetchSpinner);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Customer not found. Check the ID and try again.' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load this customer. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
      return;
    }

    const name = await ctx.promptInput(`Name (current: ${existingInfo.name})`);
    const phone = await ctx.promptInput(`Phone (current: ${existingInfo.phone_number || 'none'})`);

    const finalName = name.trim() || existingInfo.name;
    const finalPhone = phone.trim() || (existingInfo.phone_number?.toString() || '');

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Updating customer…' });
    try {
      const updated = await client.customers.update(customer_id, {
        name: finalName,
        phone_number: finalPhone !== '' ? finalPhone : null,
      });

      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Customer updated.' });
      ctx.addBlock({
        type: 'detail',
        data: {
          customer_id: updated.customer_id,
          name: updated.name,
          email: updated.email,
          phone_number: updated.phone_number,
        },
      });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't update customer. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
