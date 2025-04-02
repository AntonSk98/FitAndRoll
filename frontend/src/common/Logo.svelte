<script>
    import logo from "../assets/images/logo.png";

    import { onMount } from "svelte";
    import { getLocale, setLocale } from "./i18n";

    let isActive = false;
    let closeDelay;
    let currentLocale;

    onMount(() => (currentLocale = getLocale()));

    function changeLocale() {
        if (currentLocale === "de") {
            currentLocale = "en";
            setLocale(currentLocale);
            return;
        }

        if (currentLocale === "en") {
            currentLocale = "de";
            setLocale(currentLocale);
        }
    }
</script>

<div
    class="logo {isActive
        ? ''
        : 'hover:scale-105 transition-transform duration-500 '}"
    on:click={() => (isActive ? changeLocale() : (isActive = !isActive))}
    on:mouseover={() => clearTimeout(closeDelay)}
    on:mouseleave={() =>
        (closeDelay = setTimeout(() => (isActive = false), 600))}
>
    <img class="front-logo shadow-2xl" src={logo} class:is-active={isActive} />
    <div class="back-logo" class:is-active={isActive}>
        {#if currentLocale === "de"}
            🇩🇪
        {/if}
        {#if currentLocale === "en"}
            🇺🇸
        {/if}
    </div>
</div>

<style>
    .logo {
        position: relative;
        width: 6rem;
        height: 6rem;
        cursor: pointer;
        border-radius: 50%;
    }

    .front-logo,
    .back-logo {
        position: absolute;
        width: 6rem;
        height: 6rem;
        display: flex;
        justify-content: center;
        align-items: center;
        border-radius: 50%;
        transition:
            transform 1s ease-in-out,
            opacity 1s ease-in-out;
    }

    .front-logo {
        background: transparent;
    }

    .back-logo {
        background: var(--primary-color-dark);
        color: white;
        opacity: 0;
        transform: rotateY(180deg);
        font-size: 1.5rem;
    }

    /* Toggle effect */
    .is-active.front-logo {
        transform: rotateY(180deg);
        opacity: 0;
    }

    .is-active.back-logo {
        transform: rotateY(0deg);
        opacity: 1;
    }
</style>
