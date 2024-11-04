<!-- src/routes/+layout.svelte -->

<!-- This will contain all shared layouts across the site -->

<script>
    // Shared logic across pages goes here
    
</script>

<!-- SIDEBAR -->
<div class="dashboard-layout">
    <div class="sidebar">
        <h2 class="sidebar-title">Student</h2>
        <nav class="sidebar-items">
            <p class="selected-sidebar-item">View Tutors</p>
            <p>Appointments</p>
            <p>Messages</p>
            <p>Payments</p>
        </nav>
    </div>

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

    .sidebar {
        
        top: 0;
        bottom: 0;
        left: 0;

        padding-right: 20px;
        

        border-top-right-radius: 25px;
        border-bottom-right-radius: 25px;

        min-width: 375px;
        width: auto;

        display: flex;
        flex-direction: column;

        background-color: #453a69;
    }
    .sidebar-title {
        display: flex;
        flex-direction: column;
        
        position: relative;
        top: 0;
        left: 0;

        padding-top: 0;
        padding-left: 15px;
        margin-top: 15px;

        max-height: 50px;
        margin-bottom: 0;

        font-size: 70px;
        font-weight: bold;
        /* border: 2px solid blue */
    }
    .sidebar-items {
        display: flex;
        flex-direction: column;
        
        padding-left: 0px;
        padding-top: 50px;
        padding-bottom: 0px;

        margin: 0;

        position: relative;
        top: 0;

        font-size: 40px;
        font-weight: bold;
    }
    .sidebar-items p {
        margin: 0;
        padding: 15px 20px;
    }

    .selected-sidebar-item {
        background-color: #7261a8;
        padding-left: 20px;

        border-bottom-right-radius: 20px;
        border-top-right-radius: 20px;
    }

    .content {
        padding-left: 30px;
        padding-top: 20px
    }
    /* .page-title {
        padding-left: 30px;
    } */
</style>