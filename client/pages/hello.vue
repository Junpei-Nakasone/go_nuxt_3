<template>
    <div class="container">
        <div>
        <h3 class="title">Todos</h3>
        <p v-for="msg in createdMsg" :key="msg.id" >
            Todo:  {{ msg.name }}
            <button @click="editTodo(msg.id)">Edit</button>
            <button @click="deleteTodo(msg.id)">Delete</button>
        </p>
        </div>
        
        <hr>
        <div>
            <br>
        <!-- <button @click="getMessage()">Hello</button>
        <h2>{{ msg }}</h2> -->
        <!-- <hr> -->
        <h2 class="subtitle">Add Todo</h2>
        <form v-on:submit="newTodo">
            Todo: <input type="text" v-model="name">
            <button type="submit">Add</button>

        </form>
        </div>
    </div>
</template>

<script>
export default {
    async asyncData({ app }) {
        const res = await app.$axios.get("http://localhost:8080/todos")
        return {
            createdMsg: res.data.todos
        }
    },
    data: function() {
        return {
            msg: "",
            createdMsg: "",
            name: ""
        }
    },
    // created: {
        
    // },
    methods: {
        getMessage() {
            const res = this.$axios.get("http://localhost:8080/hello").then(res => this.msg = res.data)
        },
        newTodo() {
            this.$axios.post('http://localhost:8080/newTodo', {
                name: this.name
            }),
            this.name = '';
        },
        editTodo(id) {
            var newTodo = window.prompt("edit to this?");
            this.$axios.post('http://localhost:8080/editTodo', {
                id: id,
                new_todo: newTodo
            })
            this.reloadPage()
        },
        deleteTodo(id) {
            var check = confirm('really?');
            if (check === true) {
                this.$axios.post('http://localhost:8080/deleteTodo', {
                    id: id
                })
            }
            this.reloadPage()
        },
        reloadPage() {
            window.location.reload()
        }
    }
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
