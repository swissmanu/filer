<script lang="ts">
  import type { Rule } from "./types/rule";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher<{ apply: Rule }>();

  export let rules: ReadonlyArray<Rule>;

  const onClickRule = (r: Rule) => {
    dispatch("apply", r);
  };
</script>

<ul>
  {#each rules as rule}
    <li><button on:click={() => onClickRule(rule)}>{rule.name}</button></li>
  {/each}
</ul>
