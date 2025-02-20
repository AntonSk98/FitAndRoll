<script>
  import { t } from "svelte-i18n";

  import { writable } from "svelte/store";
  import Navbar from "./common/Navbar.svelte";
  import Courses from "./courses/Courses.svelte";
  import Toasts from "./toast/Toasts.svelte";
  import { isLoading } from "svelte-i18n";
  import "./lib/i18n.js";

  let activePage = writable($t("navigation.courses"));
</script>

{#if !$isLoading}
  <div class="my-5 mx-[3rem]">
    <Navbar bind:activePage />

    {#if $activePage === $t("navigation.courses")}
      {#key $activePage}
        <Courses />
      {/key}
    {:else if $activePage === $t("navigation.participants")}
      <div>participants</div>
    {:else if $activePage === $t("navigation.archive")}
      <div>archive</div>
    {/if}
  </div>

  <Toasts />
{/if}
