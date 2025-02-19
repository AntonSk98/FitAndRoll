<script>
  import { writable } from "svelte/store";
  import Navbar from "./common/Navbar.svelte";
  import Courses from "./courses/Courses.svelte";
    import Toasts from "./toast/Toasts.svelte";

  let activePage = writable("COURSES");
</script>


<div class="my-5 mx-[3rem]">
  <Navbar bind:activePage />

  {#if $activePage === "COURSES"}
    {#key $activePage}
      <Courses/>
    {/key}
  {:else if $activePage === "PARTICIPANTS"}
    <div>participants</div>
  {:else if $activePage === "ARCHIVE"}
    <div>archive</div>
  {/if}
</div>

<Toasts />
