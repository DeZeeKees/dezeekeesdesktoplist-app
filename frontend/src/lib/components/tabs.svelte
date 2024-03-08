<nav>

    <button class="hamburger">
        <span class="material-symbols-outlined">menu</span>
        <div class="menu">

            <div class="top">
                <a href="/" class={$page.url.pathname === '/' ? 'active' : ''}>Your List</a>
                <a href="/search" class={$page.url.pathname.includes("/search") ? 'active' : ''}>Search</a>
                <!-- @todo implement these pages -->
                <a href="/seasonal" class={$page.url.pathname === '/seasonal' ? 'active' : ''}>Seasonal</a>
                <!-- <a href="#temp" class={$page.url.pathname === '/ranking' ? 'active' : ''}>Ranking</a> -->
            </div>

            <div class="bottom">
                <a href="/settings" class={$page.url.pathname === '/settings' ? 'active' : ''}>Settings</a>
            </div>

        </div>
    </button>
    
    {#if renderTabs}
        <div class="tabs" bind:this={tabList}>
            {#each tabs as tab, i}
                    <button class={"tab" + (i === 0 ? " active" : "")} on:click={changeTab}>{tab.name}</button>
            {/each}
        </div>
    {:else}
        <div class="notabs">
            <slot />
        </div>
    {/if}

</nav>

<script>
    import { page } from "$app/stores";
    import { createEventDispatcher } from "svelte";
    export let tabs = [];
    export let renderTabs = true;

    let tabList
    const dispatch = createEventDispatcher()

    /**
     * 
     * @param event {MouseEvent}
     */
    function changeTab(event) {
        tabList.childNodes.forEach((node) => {
            node.classList.remove('active')
        })
        // @ts-ignore
        event.target.classList.add('active')

        // @ts-ignore
        dispatch("tabChange", event.target.innerText.toLowerCase())
    }
</script>

<style lang="less">

    nav {
        width: 100%;
        padding: 0.25rem;
        display: flex;
        height: 2.5rem;

        button.hamburger {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 4rem;
            height: 100%;
            border-radius: 1rem;
            background: var(--mal-blue);
            color: white;
            border: none;
            cursor: pointer;
            transition: all 0.2s ease-in-out;

            .menu {
                position: absolute;
                top: 0;
                left: -12rem;
                width: 10rem;
                height: calc(100vh - var(--titlebar-height));
                box-shadow: 0 0 1rem rgba(0, 0, 0, 0.5);
                background-color: #fff;
                z-index: 10;
                pointer-events: none;
                user-select: none;
                transition: all 0.2s ease-in-out;
                padding-inline: 0.5rem;
                padding-block: 0.5rem;

                display: flex;
                flex-direction: column;
                justify-content: space-between;

                .top {
                    display: flex;
                    flex-direction: column;
                    gap: 0.5rem;
                }

                a {
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    padding: 0.5rem;
                    border-radius: 1rem;
                    color: var(--text-dark);
                    text-decoration: none;
                    transition: all 0.2s ease-in-out;

                    &:hover {
                        background-color: var(--mal-blue);
                        color: white;
                    }

                    &.active {
                        background-color: var(--mal-blue);
                        color: white;
                    }
                }
            }
        }

        button:hover {
            .menu {
                pointer-events: all;
                left: 0;
            }
        }

        .tabs {
            width: calc(100% - 4rem);
            margin-right: 4rem;
            display: flex;
            flex-direction: row;
            justify-content: center;
            align-items: center;
            flex-grow: 1;
            padding: 0 1rem;
            overflow-x: auto;
            overflow-y: hidden;
            scrollbar-width: none;
            gap: 0.5rem;

            button {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 100%;
                padding: 0.25rem 1rem;
                border-radius: 1rem;
                background: var(--mal-blue);
                color: white;
                border: none;
                cursor: pointer;
                transition: all 0.2s ease-in-out;
            }

            button.tab.active, button.tab:hover {
                background: color-mix(in srgb, black 30%, var(--mal-blue));
            }
        }

        .notabs {
            width: calc(100% - 4rem);
        }
    }

</style>