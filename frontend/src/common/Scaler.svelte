<script>
    import { toastSuccess } from "../toast/toastStore";
    import { onDestroy, onMount } from "svelte";
    import {IS_HOTKEY_DOWN_FUNCTION, IS_HOTKEY_UP_FUNCTION} from "./scale_hotkeys";


    const defaultFontSizePx = 16;
    const stepPercent = 5;
    const baseFontStep = (defaultFontSizePx * stepPercent) / 100;
    const minFontSize = (defaultFontSizePx * 50) / 100;
    const maxFontSize = (defaultFontSizePx * 150) / 100;

    let rootFontSize;

    function persistRootFontSize(newRootFontSize) {}

    function fetchRootFontSize() {
        return null;
    }

    function handleKeyDown($event) {
        let currentSize = getCurrentRootFontSize();
        if (IS_HOTKEY_UP_FUNCTION($event)) {
            $event.preventDefault();
            const newSize = Math.min(currentSize + baseFontStep, maxFontSize);
            setRootFontSizeIfChanged(currentSize, newSize);
        } else if (IS_HOTKEY_DOWN_FUNCTION($event)) {
            $event.preventDefault();
            const newSize = Math.max(currentSize - baseFontStep, minFontSize);
            setRootFontSizeIfChanged(currentSize, newSize);
        }
    }

    function getCurrentRootFontSize() {
        if (!rootFontSize) {
            rootFontSize = fetchRootFontSize() ?? 16;
            return rootFontSize;
        }
        return parseFloat(getComputedStyle(document.documentElement).fontSize);
    }

    function setRootFontSizeIfChanged(currentPx, newPx) {
        if (currentPx === newPx) {
            return;
        }

        document.documentElement.style.fontSize = `${newPx}px`;
        const scalePercentage = ((newPx / defaultFontSizePx) * 100).toFixed(0);
        persistRootFontSize(newPx);
        toastSuccess(`${scalePercentage}%`);
    }

    onMount(() => {
        window.addEventListener("keydown", handleKeyDown);
    });

    onDestroy(() => {
        window.removeEventListener("keydown", handleKeyDown);
    });
</script>
