function changePage(pageId) {
  console.log(pageId)
    // Hide all pages
    document.querySelectorAll('.page').forEach(page => {
        page.classList.remove('active');
    });
    
    // Show selected page
    document.getElementById(pageId).classList.add('active');
    
    // Update navigation active state
    document.querySelectorAll('.nav-item').forEach(item => {
        item.classList.remove('active');
    });
    
    // Find and activate the clicked navigation item
    const navItems = document.querySelectorAll('.nav-item');
    navItems.forEach(item => {
        if (item.getAttribute('onclick').includes(pageId)) {
            item.classList.add('active');
        }
    });
}

// Initialize with menu page
document.addEventListener('DOMContentLoaded', () => {
    changePage('menuPage');
});
