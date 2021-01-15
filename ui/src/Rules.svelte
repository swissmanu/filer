<script lang="ts">
  import type { Rule } from "./types/rule";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher<{ apply: Rule; delete: void }>();

  export let rules: ReadonlyArray<Rule>;
  export let disabled: boolean = false;

  const onClickRule = (r: Rule) => {
    dispatch("apply", r);
  };
  const onClickDelete = () => {
    dispatch("delete");
  };
</script>

<ul class="flex-1">
  {#each rules as rule}
    <li class="pb-4">
      <button
        class="inline w-full px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        on:click={() => onClickRule(rule)}
        {disabled}>
        <span class="flex flex-col items-start">
          <span>{rule.name}</span>
          <span class="text-xs font-normal">{rule.description}</span>
        </span>
      </button>
    </li>
  {/each}
  <li class="pb-4">
    <button
      class="inline w-full px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-500 hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      on:click={onClickDelete}>
      <span class="flex flex-col items-start text-left">
        <span>Löschen</span>
        <span class="text-xs font-normal"
          >Dieses Dokument unwiederruflich löschen</span
        >
      </span>
    </button>
  </li>
</ul>
