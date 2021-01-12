<script lang="ts">
  import { onMount } from "svelte";
  import { API } from "./api";
  import type { InboxItem } from "./types/inboxItem";
  import type { Rule } from "./types/rule";
  import PdfViewer from "./PdfViewer.svelte";
  import Rules from "./Rules.svelte";

  const api = new API();

  let rules: ReadonlyArray<Rule>;
  let inboxItems: ReadonlyArray<InboxItem>;

  onMount(async () => {
    const [rs, is] = await Promise.all([api.getRules(), api.getInbox()]);
    rules = rs;
    inboxItems = is;
  });

  const onApplyRule = async (rule: Rule) => {
    try {
      await api.applyRuleToInboxItem(rule, inboxItems[0]);
      inboxItems = inboxItems.slice(1);
    } catch (e) {}
  };
</script>

<main class="bg-red-100">
  <h1>filer</h1>
  {#if !inboxItems || !rules}
    <div>Loading...</div>
  {:else if inboxItems[0]}
    <section>
      <header>
        {inboxItems.length} documents pending
      </header>
      <PdfViewer url={'/inbox/' + inboxItems[0].name} />
      <Rules {rules} on:apply={({ detail: rule }) => onApplyRule(rule)} />
    </section>
  {:else}
    <div>Nothing to file! 😎</div>
  {/if}
</main>
