<script lang="ts">
  import { onMount } from "svelte";
  import { API } from "./api";
  import type { InboxItem } from "./types/inboxItem";
  import type { Rule } from "./types/rule";

  let rules: ReadonlyArray<Rule> | null = null;
  let inboxItems: ReadonlyArray<InboxItem> | null = null;

  onMount(async () => {
    const api = new API("http://localhost:8000");
    rules = await api.getRules();
    inboxItems = await api.getInbox();
  });
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
  <h1>filer</h1>
  <p>{JSON.stringify(rules)}</p>
  <p>{JSON.stringify(inboxItems)}</p>
</main>
