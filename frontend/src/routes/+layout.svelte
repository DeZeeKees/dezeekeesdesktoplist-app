<svelte:head>
    <link rel="stylesheet" href="fonts/fonts.css">
</svelte:head>

<svelte:window on:keydown={GlobalKeyDown} />

<script>
    import Titlebar from "$lib/components/titlebar.svelte";
    import { EventsEmit, EventsOn } from '$lib/wailsjs/runtime'
    import { goto } from '$app/navigation'
    import { popover, title } from "$lib/store";
    import { page } from "$app/stores";

    let showTitlebar = true

    EventsOn('svelte:goto', (data) => {
        goto(data)
    })

    EventsEmit('startup:success')

    function GlobalKeyDown(event) {
        if(
            $popover.element !== undefined &&
            $page.url.pathname === '/'&&
            $popover.element.classList.contains("inactive") === false
        ) {
            switch(event.key) {
                case "Escape":
                    $popover.close()
                    break;
                case "Enter":
                    $popover.save()
                    break;
            }
        }

        if(event.ctrlKey && event.key === "F11") {
            event.preventDefault()
            showTitlebar = !showTitlebar

            if(showTitlebar) {
                document.documentElement.style.setProperty('--titlebar-height', '2rem');
            } else {
                document.documentElement.style.setProperty('--titlebar-height', '0rem');
            }
        }
    }
</script>

{#if showTitlebar}
    <Titlebar title={$title} />
{/if}

<div id="wrapper" class="wrapper">
    <slot/>
</div>

<style lang="less">

:global(*) {
    box-sizing: border-box;
    font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}

:global(body), :global(html) {
    margin: 0;
    padding: 0;
}

:global(:root) {
    --titlebar-height: 2rem;
    --your-list-card-size-multiplier: 1;

    --mal-blue: #2E51A2;
    --mal-blue-dark: color-mix(in srgb, black 10%, var(--mal-blue));

    --text-dark: #455A64;
    --text-light: #CFD8DC;
    --text-white: #FFFFFF;
    --text-black: #000000;
}
  
:global(.material-symbols-outlined) {
    font-family: 'Material Symbols Outlined';
    font-weight: normal;
    font-style: normal;
    font-size: 24px;
    line-height: 1;
    letter-spacing: normal;
    text-transform: none;
    display: inline-block;
    white-space: nowrap;
    word-wrap: normal;
    direction: ltr;
    -webkit-font-smoothing: antialiased;
}

.wrapper {
    position: relative;
    height: calc(100vh - var(--titlebar-height));
    overflow: hidden;
}

:global(.toast-down) {
    margin-top: 2rem;
}

</style>