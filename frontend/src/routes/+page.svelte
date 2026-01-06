<script>
	import { GetTodoList, AddTodo, DeleteTodo, UpdateTodo } from '$lib/wailsjs/go/main/WailsApp.js';

	let todoList = $state([]);

	let todoData = $state('');

	$effect(async () => {
		todoList = await GetTodoList();
	});

	const addTodo = async () => {
		if (todoData == '' || todoData == null) {
			window.alert('No Data Entered');
			return;
		}
		await AddTodo(todoData);
		todoList = await GetTodoList();
		todoData = '';
	};

	const deleteTodo = async (id) => {
		await DeleteTodo(id);
		todoList = await GetTodoList();
	};

	const updateTodo = async (id, state) => {
		await UpdateTodo(id, state);
		todoList = await GetTodoList();
	};

	const formatDate = (date) => {
		const months = [
			'January',
			'February',
			'March',
			'April',
			'May',
			'June',
			'July',
			'August',
			'September',
			'October',
			'November',
			'December'
		];

		const valDate = new Date(date);
		return {
			date: `${valDate.getDate()} ${months[valDate.getMonth()]}`,
			time: valDate.toLocaleTimeString('en-US', {
				hour: 'numeric',
				minute: 'numeric',
				hour12: true
			})
		};
	};
</script>

<main class="flex h-full min-h-screen w-full flex-col gap-8 bg-slate-900 p-4">
	<div class="mx-auto flex h-fit w-full max-w-2xl flex-row items-center justify-center gap-4">
		<input
			type="text"
			class="h-12 w-full rounded-xl bg-slate-950 px-2 text-xl text-slate-50 outline-0 focus-within:outline-2 focus-within:outline-slate-400 focus:outline-slate-600"
			bind:value={todoData}
		/>
		<button
			class="h-12 w-fit flex-none rounded-xl bg-green-600 px-6 text-xl font-bold text-green-50 active:scale-95 active:bg-green-800"
			onclick={addTodo}
		>
			Add
		</button>
	</div>

	{#snippet item(todo)}
		{@const { date, time } = formatDate(todo.updated)}

		<div
			class={[
				'flex h-fit w-full flex-row items-center justify-start gap-4 overflow-hidden rounded-xl pr-4',
				todo.state ? 'bg-slate-800' : 'bg-slate-950'
			]}
		>
			<label
				class={[
					'flex h-full w-full flex-row gap-4 overflow-hidden p-4 text-nowrap text-ellipsis',
					todo.state ? 'text-slate-600' : 'text-slate-100'
				]}
			>
				<input
					type="checkbox"
					class="hidden"
					bind:checked={todo.state}
					onchange={() => updateTodo(todo.id, todo.state)}
				/>
				<div
					class={[
						'flex h-fit w-fit flex-none flex-col items-start justify-center',
						todo.state ? 'text-slate-600' : 'text-slate-400'
					]}
				>
					<span class="text-xs">{date}</span>
					<span class="text-xs">{time}</span>
				</div>
				{todo.data}
			</label>

			<button
				class="size-8 flex-none rounded-xl bg-red-400 text-2xl font-bold active:scale-90 active:bg-red-600"
				onclick={() => deleteTodo(todo.id)}
			>
				X
			</button>
		</div>
	{/snippet}

	<div class="mx-auto flex h-fit w-full max-w-2xl flex-col gap-4">
		{#each todoList as todo (todo.id)}
			{@render item(todo)}
		{/each}
	</div>
</main>
