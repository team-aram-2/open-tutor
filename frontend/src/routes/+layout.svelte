<!-- src/routes/+layout.svelte -->

<!-- This will contain all shared layouts across the site -->

<script>
    // Shared logic across pages goes here
    import '../app.css';
    import Sidebar from '$lib/components/sidebar.svelte';

    // Load student my people dashboard by default TEMPORARY
    import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

    onMount(() => {
        goto('/my_people/student');
    });
</script>



<!-- Page Content -->
<div class="dashboard-layout">

    <Sidebar>
    </Sidebar>
    

    <main class="content">


        <slot />
        <!-- Where page content will be rendered -->
         
    </main>
</div>

<!-- Style for the page goes here -->
<style>
    /* Import Inter font */
    @import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700');

    /* Apply these globally */
    * {
        font-family: 'Inter', sans-serif;
    }

    .dashboard-layout {
        display: flex;
        flex-direction: row;
        position: fixed;

        top: 0;
        bottom: 0;
        left: 0;
        right: 0;

        min-height: 100;

        background: #231d34;

        /* font:  */
    }

    .content {
        padding-left: 30px;
        padding-top: 20px
    }
    /* .page-title {
        padding-left: 30px;
    } */
</style>