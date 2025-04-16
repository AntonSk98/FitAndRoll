<script>
  import { writable } from "svelte/store";
  import Navbar from "./common/Navbar.svelte";
  import Courses from "./courses/Courses.svelte";
  import Toasts from "./toast/Toasts.svelte";
  import { isLoading, locale } from "svelte-i18n";
  import "./lib/i18n.js";
  import Participants from "./participants/Participants.svelte";
  import Archive from "./archive/Archive.svelte";
  import Statistics from "./statistics/Statistics.svelte";
  import Logo from "./common/Logo.svelte";
  import Scaler from "./common/Scaler.svelte";

  let activePage = writable("courses");

  let key = 0;

  $: {
    if ($locale) {
      key++;
    }
  }
</script>

{#if !$isLoading}
  <div class="my-5 mx-[3rem]">
    <div class="flex gap-3">
      <Logo />
      <Navbar bind:activePage localeSet={key} />
    </div>
    {#key key}
      {#if $activePage === "courses"}
        <Courses />
      {:else if $activePage === "participants"}
        <Participants />
      {:else if $activePage === "archive"}
        <Archive />
      {:else if $activePage === "statistics"}
        <Statistics />
      {/if}
    {/key}
  </div>

  <Toasts />
  <Scaler />
{/if}
