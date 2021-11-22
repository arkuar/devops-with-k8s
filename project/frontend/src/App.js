import axios from 'axios';
import React, {useEffect, useState} from 'react';
import './App.css';

const APP_BACKEND_URL = process.env.APP_BACKEND_URL || '/api'

const imageUrl = `${APP_BACKEND_URL}/image`

async function fetchTodos() {
  const result = await axios.get(`${APP_BACKEND_URL}/todos`)
  return result?.data?.todos ?? []
}

function App() {
  const [todos, setTodos] = useState([])
  const [todoBody, setTodoBody] = useState("")

  useEffect(() => {
    (async function(){
      try {
        setTodos(await fetchTodos())
      } catch (error) {
        console.error("Error fetching todos", error)
      }
    })()
  }, [])

  const createTodo = async () => {
    try {
      const result = await axios.post(`${APP_BACKEND_URL}/todos`, {todo: todoBody})
      setTodoBody("")
      setTodos((current) => current.concat(result.data.todo))
    } catch (error) {
    }
  }

  const handleInputChange = (e) => setTodoBody(e.target.value) 

  return (
    <div className="App">
      <h1>Hello from Kubernetes</h1>
      <img src={imageUrl} width="200" height="200" alt="img" />
      <input type="text" value={todoBody} onChange={handleInputChange} maxlength="140" />
      <button type="button" onClick={createTodo}>Create TODO</button>
      <ul>
        {todos.map((todo) => <li key={todo}>{todo}</li>)}
      </ul>
    </div>
  );
}

export default App;
