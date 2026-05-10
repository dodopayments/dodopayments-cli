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

export const renderBlock = (block: BlockType) => (
  <Switch fallback={null}>
    <Match when={block.type === 'spinner' && block}>
      {(b) => <Spinner label={(b() as Extract<BlockType, { type: 'spinner' }>).label} />}
    </Match>
    <Match when={block.type === 'success' && block}>
      {(b) => <Success message={(b() as Extract<BlockType, { type: 'success' }>).message} />}
    </Match>
    <Match when={block.type === 'error' && block}>
      {(b) => <Error message={(b() as Extract<BlockType, { type: 'error' }>).message} />}
    </Match>
    <Match when={block.type === 'info' && block}>
      {(b) => <Info message={(b() as Extract<BlockType, { type: 'info' }>).message} />}
    </Match>
    <Match when={block.type === 'empty' && block}>{() => <Empty />}</Match>
    <Match when={block.type === 'link' && block}>
      {(b) => {
        const blk = b() as Extract<BlockType, { type: 'link' }>;
        return <Link text={blk.text} url={blk.url} />;
      }}
    </Match>
    <Match when={block.type === 'event' && block}>
      {(b) => <Event event={(b() as Extract<BlockType, { type: 'event' }>).event} />}
    </Match>
    <Match when={block.type === 'streaming' && block}>
      {(b) => <Streaming text={(b() as Extract<BlockType, { type: 'streaming' }>).text} />}
    </Match>
    <Match when={block.type === 'table' && block}>
      {(b) => {
        const blk = b() as Extract<BlockType, { type: 'table' }>;
        return <Table data={blk.data} statusColumn={blk.statusColumn} />;
      }}
    </Match>
    <Match when={block.type === 'detail' && block}>
      {(b) => <Detail data={(b() as Extract<BlockType, { type: 'detail' }>).data} />}
    </Match>
    <Match when={block.type === 'help' && block}>{() => <Help />}</Match>
    <Match when={block.type === 'inline-input' && block}>
      {(b) => {
        const blk = b() as Extract<BlockType, { type: 'inline-input' }>;
        return <InlineInput label={blk.label} secure={blk.secure} onSubmit={blk.onSubmit} />;
      }}
    </Match>
    <Match when={block.type === 'inline-select' && block}>
      {(b) => {
        const blk = b() as Extract<BlockType, { type: 'inline-select' }>;
        return (
          <InlineSelect label={blk.label} options={blk.options} onSubmit={blk.onSubmit} />
        );
      }}
    </Match>
    <Match when={block.type === 'confirm' && block}>
      {(b) => {
        const blk = b() as Extract<BlockType, { type: 'confirm' }>;
        return (
          <Confirm
            message={blk.message}
            onConfirm={blk.onConfirm}
            onCancel={blk.onCancel}
          />
        );
      }}
    </Match>
  </Switch>
);
