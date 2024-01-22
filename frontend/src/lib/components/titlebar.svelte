<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="titlebar" on:dblclick={WindowToggleMaximise}>

    <div class="titlebar__title">
        <h1>{title}</h1>
    </div>

    <div class="titlebar__actions">
        <button class="titlebar__action" on:mousedown={WindowMinimise}>
            <span class="material-symbols-outlined">minimize</span>
        </button>
        <button class="titlebar__action" on:mousedown={WindowToggleMaximise}>
            <span class="material-symbols-outlined">fullscreen</span>
        </button>
        <button class="titlebar__action" on:mousedown={Quit}>
            <span class="material-symbols-outlined">close</span>
        </button>
    </div>

</div>

<script>
    import { WindowMinimise, Quit, WindowToggleMaximise } from "$lib/wailsjs/runtime";

    export let title = "";
</script>

<style lang="less">

    .titlebar {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        background-color: #2c3e50;
        color: #ecf0f1;
        height: var(--titlebar-height);
        --wails-draggable: drag;
        user-select: none;
        -webkit-app-region: drag;
    }

    .titlebar__title {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
        padding-left: 1rem;
    }

    .titlebar__title h1 {
        font-size: 1rem;
        margin: 0;
        padding: 0;
    }

    .titlebar__actions {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
        height: 100%;
    }

    .titlebar__action {
        background-color: transparent;
        border: none;
        color: #ecf0f1;
        font-size: 1rem;
        margin: 0;
        padding: 0;
        cursor: pointer;
        -webkit-app-region: no-drag;
        position: relative;
        height: 100%;
        aspect-ratio: 1/1;

        .material-symbols-outlined {
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
        }
    }

    .titlebar__action:hover {
        color: #bdc3c7;
    }

    .titlebar__action:active {
        color: #95a5a6;
    }

</style>