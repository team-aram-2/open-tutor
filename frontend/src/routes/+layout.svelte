<!-- src/routes/+layout.svelte -->

<!-- This will contain all shared layouts across the site -->

<script>
    // Shared logic across pages goes here
    import '../app.css';
    import Sidebar from '$lib/components/sidebar.svelte';
    
</script>



<!-- SIDEBAR -->
<div class="dashboard-layout">

    <Sidebar>
        <h1 class="text-4xl font-bold">Welcome to My App!</h1>
        <p class="text-lg mt-4">This is a simple demo of a collapsible sidebar with SvelteKit and TailwindCSS.</p>
    </Sidebar>
    

    <main class="content">
        <header>
            <h1>DASHBOARD HEADER PLACEHOLDER</h1>
        </header>

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