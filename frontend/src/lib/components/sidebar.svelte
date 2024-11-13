<script>
  let selectedItem = 'view-tutors';

  const setSelectedItem = (item) => {
    selectedItem = item;
  };
</script>

<div class="sidebar flex h-screen">
    <h2 class="sidebar-title">Student</h2>
    <nav class="sidebar-items">
        <a href="#" class="no-decoration" class:selected-sidebar-item={selectedItem === 'view'} on:click={() => setSelectedItem('view')}>
            <p>View Tutors</p>
        </a>
        <a href="#" class="no-decoration" class:selected-sidebar-item={selectedItem === 'apt'} on:click={() => setSelectedItem('apt')}>
            <p>Appointments</p>
        </a>
        <a href="#" class="no-decoration" class:selected-sidebar-item={selectedItem === 'msg'} on:click={() => setSelectedItem('msg')}>
            <p>Messages</p>
        </a>
        <a href="#" class="no-decoration" class:selected-sidebar-item={selectedItem === 'pym'} on:click={() => setSelectedItem('pym')}>
            <p>Payments</p>
        </a>
    </nav>
</div>

<style>
    
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
        color: white;
        /* border: 2px solid blue */
    }
    .sidebar-items {
        display: flex;
        flex-direction: column;
        
        padding-left: 0px;
        padding-top: 50px;
        padding-bottom: 50px;

        margin: 0;

        position: relative;
        top: 0;

        font-size: 40px;
        font-weight: bold;

        
    }
    .sidebar-items p {
        margin: 0;
        padding: 15px 20px;
        color: white;
    }

    .selected-sidebar-item {
        background-color: #7261a8;
        padding-left: 20px;

        border-bottom-right-radius: 20px;
        border-top-right-radius: 20px;
    }
    .no-decoration{
        text-decoration: none;
        color: inherit;
    }
</style>