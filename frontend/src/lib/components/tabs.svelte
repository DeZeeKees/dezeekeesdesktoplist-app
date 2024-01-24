<nav>

    <button class="hamburger">
        <span class="material-symbols-outlined">menu</span>
        <div class="menu">

            <div class="top">
                <a href="/" class={$page.url.pathname === '/' ? 'active' : ''}>Your List</a>
                <a href="/seasonal" class={$page.url.pathname === '/seasonal' ? 'active' : ''}>Seasonal</a>
                <a href="/ranking" class={$page.url.pathname === '/ranking' ? 'active' : ''}>Ranking</a>
            </div>

            <div class="bottom">
                <a href="/settings" class={$page.url.pathname === '/settings' ? 'active' : ''}>Settings</a>
            </div>

        </div>
    </button>
    
    <div class="tabs" bind:this={tabList}>
        {#each tabs as tab, i}
                <button class={"tab" + (i === 0 ? " active" : "")} on:click={changeTab}>{tab.name}</button>
        {/each}
    </div>

</nav>

<script>
    import { page } from "$app/stores";
    import { createEventDispatcher } from "svelte";
    export let tabs = [];

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
                left: -10rem;
                width: 10rem;
                height: calc(100vh - var(--titlebar-height));
                border-right: solid 0.1rem var(--text-dark);
                background-color: #fff;
                z-index: 10;
                pointer-events: none;
                user-select: none;
                transition: all 0.2s ease-in-out;

                display: flex;
                flex-direction: column;
                justify-content: space-between;

                a {
                    width: 100%;
                    height: 2.5rem;
                    padding: 0 1rem;
                    display: flex;
                    justify-content: flex-start;
                    align-items: center;
                    color: var(--text-dark);
                    text-decoration: none;
                    transition: all 0.2s ease-in-out;

                    &:hover, &.active {
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
    }

</style>