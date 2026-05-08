import DodoPayments from 'dodopayments';
import chalk from 'chalk';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleCustomers(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching customers...' });
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
      ctx.addBlock({ type: 'streaming', text: chalk.dim('\nTip: Use ') + chalk.cyan('/customers list <page>') + chalk.dim(' to see more pages.') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch customers: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'create') {
    const name = await ctx.promptInput('Enter Name:');
    const email = await ctx.promptInput('Enter Email:');
    const phone = await ctx.promptInput('Enter Phone Number:');
    
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Creating customer...' });
    try {
      const creation = await client.customers.create({
        name,
        email,
        phone_number: phone.trim() !== '' ? phone : null,
      });

      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Customer Successfully Created!' });
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
        ctx.addBlock({ type: 'error', message: `Failed to create customer: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'update') {
    const customer_id = args[0];
    if (!customer_id) {
      ctx.addBlock({ type: 'error', message: 'Please provide a customer ID! (Usage: /customers update <id>)' });
      return;
    }
    
    const fetchSpinner = ctx.addBlock({ type: 'spinner', label: 'Fetching customer details...' });
    let existingInfo;
    try {
      existingInfo = await client.customers.retrieve(customer_id);
      ctx.removeBlock(fetchSpinner);
    } catch (e: any) {
      ctx.removeBlock(fetchSpinner);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Incorrect customer ID!' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch customer: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
      return;
    }

    const name = await ctx.promptInput(`Enter customer name (Current: ${existingInfo.name}):`);
    const phone = await ctx.promptInput(`Enter customer phone (Current: ${existingInfo.phone_number || 'None'}):`);

    const finalName = name.trim() || existingInfo.name;
    const finalPhone = phone.trim() || (existingInfo.phone_number?.toString() || '');

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Updating customer...' });
    try {
      const updated = await client.customers.update(customer_id, {
        name: finalName,
        phone_number: finalPhone !== '' ? finalPhone : null,
      });

      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Customer Successfully Updated!' });
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
        ctx.addBlock({ type: 'error', message: `Failed to update customer: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
