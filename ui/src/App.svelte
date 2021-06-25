<script lang="ts">
  import { onMount } from "svelte";
  import { isLoading, _ } from "svelte-i18n";
  import { API } from "./api";
  import PdfViewer from "./PdfViewer.svelte";
  import Rules from "./Rules.svelte";
  import type { InboxItem } from "./types/inboxItem";
  import type { Rule } from "./types/rule";

  const api = new API();

  let rules: ReadonlyArray<Rule>;
  let inboxItems: ReadonlyArray<InboxItem>;
  let name: string = "";
  $: ruleButtonsDisabled = name.trim().length === 0;

  onMount(async () => {
    const [rs, is] = await Promise.all([api.getRules(), api.getInbox()]);
    rules = rs;
    inboxItems = is;
    getNameFromFirstInboxItem();
  });

  const getNameFromFirstInboxItem = () => {
    name = inboxItems[0].name.replace(/\.pdf$/i, "");
  };

  const nextInboxItem = () => {
    inboxItems = inboxItems.slice(1);
    getNameFromFirstInboxItem();
  };

  const onApplyRule = async (rule: Rule) => {
    try {
      await api.applyRuleToInboxItem(
        rule,
        inboxItems[0],
        name !== inboxItems[0].name ? name : undefined
      );
      nextInboxItem();
    } catch (e) {}
  };

  const onDeleteDoucment = async () => {
    if (window.confirm($_("delete.confirm"))) {
      try {
        await api.deleteInboxItem(inboxItems[0]);
        nextInboxItem();
      } catch (e) {}
    }
  };
</script>

<svelte:head>
  <title>{$_("inbox.title")}</title>
</svelte:head>

<main class="max-w-7xl mx-auto sm:px-6 lg:px-8">
  {#if $isLoading}{:else}
    <section
      class="bg-white overflow-hidden shadow rounded-lg divide-y divide-gray-200 md:mt-16"
    >
      <header class="px-4 py-5 sm:px-6">
        <h1 class="text-2xl font-medium">{$_("inbox.title")}</h1>
        {#if inboxItems && inboxItems.length > 0}
          <p class="text-sm" data-testid="pending-inbox-items">
            {$_("inbox.pendingDocuments", {
              values: { count: inboxItems.length },
            })}
          </p>
        {/if}
      </header>
      <div class="px-4 py-5 sm:p-6">
        {#if !inboxItems || !rules}
          <div>{$_("inbox.loading")}</div>
        {:else if inboxItems[0]}
          <h2 class="text-lg medium mb-3">
            <input
              data-testid="inbox-item-name"
              type="text"
              name="inboxItemName"
              id="inboxItemName"
              class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
              placeholder={$_("inbox.rename.placeholder")}
              bind:value={name}
            />
          </h2>
          <div class="flex gap-4 flex-col md:flex-row">
            <div class="flex-2">
              <PdfViewer url={api.getUrlForItem(inboxItems[0])} />
            </div>
            <div class="flex-1">
              <Rules
                {rules}
                disabled={ruleButtonsDisabled}
                on:apply={({ detail: rule }) => onApplyRule(rule)}
                on:delete={onDeleteDoucment}
              />
            </div>
          </div>
        {:else}
          <p>{$_("inbox.noPendingDocuments")}</p>
        {/if}
      </div>
    </section>
  {/if}
</main>
