<script>
  import { writable } from "svelte/store";
  import Navbar from "./common/Navbar.svelte";
  import Courses from "./courses/Courses.svelte";
  import Toasts from "./toast/Toasts.svelte";
  import { isLoading } from "svelte-i18n";
  import "./lib/i18n.js";
  import Participants from "./participants/Participants.svelte";
    import Archive from "./archive/Archive.svelte";
    import Statistics from "./statistics/Statistics.svelte";

  let activePage = writable("courses");
</script>

{#if !$isLoading}
  <div class="my-5 mx-[3rem]">
    <Navbar bind:activePage />

    {#if $activePage === "courses"}
      <Courses />
    {:else if $activePage === "participants"}
      <Participants />
    {:else if $activePage === "archive"}
      <Archive/>
    {:else if $activePage === "statistics"}
      <Statistics/>
    {/if}
  </div>

  <Toasts />
{/if}
