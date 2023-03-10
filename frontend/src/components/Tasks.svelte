<script>
    import Task from "./Task.svelte";
    let tasks;
    let isLoading = true;

    fetch("http://localhost:9000/tasks", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    })
        .then(async (res) => {
            tasks = await res.json();
            isLoading = false;
        })
        .catch((err) => {
            console.log("error fetching tasks from backend: ", err);
            isLoading = false;
        });
</script>

<div
    class="mx-auto border border-gray-300 rounded-lg mt-24 max-w-screen-sm py-2.5 px-4"
>
    {#if isLoading}
        <p>Loading...</p>
    {:else if tasks}
        {#each tasks as task}
            <Task {task} />
        {/each}
    {/if}
</div>
