import axios from 'axios';
import React, {useEffect, useState} from 'react';
import './App.css';

const APP_BACKEND_URL = process.env.APP_BACKEND_URL || '/api'

const imageUrl = `${APP_BACKEND_URL}/image`

function App() {
  const [todos, setTodos] = useState([])
  const [todoContent, setTodoContent] = useState("")

  useEffect(() => {
    (async function(){
      try {
        const result = await axios.get(`${APP_BACKEND_URL}/todos`)
        setTodos(result?.data?.todos ?? [])
      } catch (error) {
        console.error("Error fetching todos", error)
      }
    })()
  }, [])

  const createTodo = async () => {
    try {
      const result = await axios.post(`${APP_BACKEND_URL}/todos`, {content: todoContent})
      setTodoContent("")
      setTodos((current) => current.concat(result.data.todo))
    } catch (error) {
    }
  }

  const handleInputChange = (e) => setTodoContent(e.target.value) 

  const markDone = async (todo) => {
    try {
      const {data: {todo: updatedTodo}} = await axios.put(`${APP_BACKEND_URL}/todos/${todo.id}`, {
        ...todo,
        done: !todo.done
      })
      setTodos((current) => current.map(t => t.id === updatedTodo.id ? updatedTodo : t))
    } catch (error) {
    }
  }

  return (
    <div className="App">
      <h1>Hello from Kubernetes</h1>
      <img src={imageUrl} width="200" height="200" alt="img" />
      <input type="text" value={todoContent} onChange={handleInputChange} maxlength="140" />
      <button type="button" onClick={createTodo}>Create TODO</button>
      <ul>
        {todos.map((todo) => <li key={todo.id} className={todo.done ? "done" : ""} onClick={() => markDone(todo)}>{todo.content}</li>)}
      </ul>
    </div>
  );
}

export default App;
