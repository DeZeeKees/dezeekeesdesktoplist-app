<main>

    <div class="card">
        <h1>MyAnimeList Login</h1>
        <p>Log in with your MyAnimeList account to sync your anime and manga lists.</p>

        <button on:mousedown={StartAuthProcess}>
            <img src="MyAnimeList_Logo.png" alt="MAL Logo">
            <span>Login With MAL</span>
        </button>
    </div>



</main>

<script>
    import { title } from '$lib/store';
    import { StartAuthProcess } from '$lib/wailsjs/go/main/App';
    import { onMount } from 'svelte';

    onMount(() => {
        title.set('Login');
    })
</script>

<style lang="less">

    main {
        display: flex;
        justify-content: center;
        align-items: center;
        height: calc(100vh - var(--titlebar-height));
    }

    .card {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 2rem;
        border-radius: 1rem;
        background: var(--mal-blue);
        color: white;
        box-shadow: 0 0 1rem rgba(0, 0, 0, 0.5);
    }

    h1 {
        font-size: 2rem;
        margin: 0;
        margin-bottom: 1rem;
    }

    p {
        font-size: 1.5rem;
        margin: 0;
        margin-bottom: 2rem;
    }

    button {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0 1rem;
        border-radius: 1rem;
        background: var(--mal-blue);
        color: white;
        border: none;
        cursor: pointer;
        transition: all 0.2s ease-in-out;
        border: solid 0.1rem var(--mal-blue-dark);

        img {
            width: 4rem;
            height: 4rem;
            margin-right: 1rem;
            filter: invert(1);
        }

        span {
            font-size: 1.5rem;
        }
    }

    button:hover {
        background: var(--mal-blue-dark);
    }

</style>