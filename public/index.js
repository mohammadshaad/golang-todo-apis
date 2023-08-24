function removeFromDb(item){
    fetch(`/delete?item=${item}`, { method: "DELETE" })
        .then(res => {
            if (res.status == 200) {
                window.location.pathname = "/";
            }
        })
        .catch(error => {
            console.error("Error deleting item:", error);
        });
}

function updateDb(oldItem, newItem) {
    console.log("Updating item:", oldItem, "with new value:", newItem);
    fetch(`/update?olditem=${oldItem}&newitem=${newItem}`, { method: "PUT" })
        .then(res => {
            if (res.status == 200) {
                alert("Database updated");
                window.location.pathname = "/";
            }
        })
        .catch(error => {
            console.error("Error updating item:", error);
        });
}

