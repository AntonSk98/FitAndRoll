<script>
    import { toastError, toastSuccess } from "../toast/toastStore";
    import { onDestroy, onMount } from "svelte";
    import { IS_HOTKEY_DOWN_FUNCTION, IS_HOTKEY_UP_FUNCTION } from "./scale_hotkeys";
    import { FetchRootFontSize, PersistRootFontSize } from "../../wailsjs/go/scaler/ScalerHandler";

    const defaultFontSizePx = 16;
    const stepPercent = 5;
    const baseFontStep = (defaultFontSizePx * stepPercent) / 100;
    const minFontSize = (defaultFontSizePx * 50) / 100;
    const maxFontSize = (defaultFontSizePx * 150) / 100;

    let rootFontSize;
    let initialized = false;

    function scaleApplication($event) {
        if (!initialized) {
            return;
        }
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
        return parseFloat(getComputedStyle(document.documentElement).fontSize);
    }

    function setRootFontSizeIfChanged(currentPx, newPx, notify=true) {
        if (currentPx === newPx) {
            return;
        }

        const effectiveRootFontSize = Math.round(newPx * 100) / 100;
        document.documentElement.style.fontSize = `${effectiveRootFontSize}px`;
        PersistRootFontSize(effectiveRootFontSize)
        .catch(err => {
            toastError();
            console.error(err);
        });

        if (notify) {
            const scalePercentage = ((effectiveRootFontSize / defaultFontSizePx) * 100).toFixed(0);
            toastSuccess(`${scalePercentage}%`);
        }
    }

    onMount(async () => {
        FetchRootFontSize().then(fetchedSize => {
            rootFontSize = fetchedSize > 0 ? fetchedSize : defaultFontSizePx;
            setRootFontSizeIfChanged(parseFloat(getComputedStyle(document.documentElement).fontSize), rootFontSize, false);
            initialized = true;
        })
        window.addEventListener("keydown", scaleApplication);
    });

    onDestroy(() => {
        window.removeEventListener("keydown", scaleApplication);
    });
</script>
