/**
 * Block dispatcher. Switches on block.type and renders the matching Solid
 * component. Adding a new BlockVariant means: extend types.ts, add a Match
 * arm here, write the renderer file. Failure to add a Match arm renders
 * nothing (silent drop) -- the dev demo (_blocks_demo) catches that.
 */

import { Match, Switch } from 'solid-js';
import type { BlockType } from '../../types';
import { Spinner } from './Spinner';
import { Success, Error, Info, Empty, Link, Event, Streaming } from './Status';
import { Table } from './Table';
import { Detail } from './Detail';
import { Help } from './Help';
import { InlineInput } from './InlineInput';
import { InlineSelect } from './InlineSelect';
import { Confirm } from './Confirm';

type Variant<T extends BlockType['type']> = Extract<BlockType, { type: T }>;

export const renderBlock = (block: BlockType) => (
  <Switch fallback={null}>
    <Match when={block.type === 'spinner' && block}>
      {(b: () => Variant<'spinner'>) => <Spinner label={b().label} />}
    </Match>
    <Match when={block.type === 'success' && block}>
      {(b: () => Variant<'success'>) => <Success message={b().message} />}
    </Match>
    <Match when={block.type === 'error' && block}>
      {(b: () => Variant<'error'>) => <Error message={b().message} />}
    </Match>
    <Match when={block.type === 'info' && block}>
      {(b: () => Variant<'info'>) => <Info message={b().message} />}
    </Match>
    <Match when={block.type === 'empty' && block}>{() => <Empty />}</Match>
    <Match when={block.type === 'link' && block}>
      {(b: () => Variant<'link'>) => <Link text={b().text} url={b().url} />}
    </Match>
    <Match when={block.type === 'event' && block}>
      {(b: () => Variant<'event'>) => <Event event={b().event} />}
    </Match>
    <Match when={block.type === 'streaming' && block}>
      {(b: () => Variant<'streaming'>) => <Streaming text={b().text} />}
    </Match>
    <Match when={block.type === 'table' && block}>
      {(b: () => Variant<'table'>) => (
        <Table data={b().data} statusColumn={b().statusColumn} />
      )}
    </Match>
    <Match when={block.type === 'detail' && block}>
      {(b: () => Variant<'detail'>) => <Detail data={b().data} />}
    </Match>
    <Match when={block.type === 'help' && block}>{() => <Help />}</Match>
    <Match when={block.type === 'inline-input' && block}>
      {(b: () => Variant<'inline-input'>) => (
        <InlineInput label={b().label} secure={b().secure} onSubmit={b().onSubmit} />
      )}
    </Match>
    <Match when={block.type === 'inline-select' && block}>
      {(b: () => Variant<'inline-select'>) => (
        <InlineSelect label={b().label} options={b().options} onSubmit={b().onSubmit} />
      )}
    </Match>
    <Match when={block.type === 'confirm' && block}>
      {(b: () => Variant<'confirm'>) => (
        <Confirm
          message={b().message}
          onConfirm={b().onConfirm}
          onCancel={b().onCancel}
        />
      )}
    </Match>
  </Switch>
);
