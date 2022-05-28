import React from 'react';
import axios from 'axios';
import './App.css';

export default function App() {

  const [users, setUsers] = React.useState<User[] | null>([]);
  const [posts, setPosts] = React.useState<Post[] | null>([]);

  let displayUsers = (users: User[] | null) => {
    if (users!==null) {
      for (const user of users) {
        return <h1>Hello user: {user.name}</h1>
      }
    }
  }
  
  async function getUsers() {
    const { data, status } = await axios.get<User[] | null>('http://localhost:8080/users')
    if (status === 200) {
      console.log(data)
      return data
    } else {
      console.log('error')
      return null
    }
  }

  async function getPosts(user: User) {
    const { data, status } = await axios.get<Post[] | null>('http://localhost:8080/users')
    if (status === 200) {
      console.log(data)
      return data
    } else {
      console.log('error')
      return null
    }
  }

  React.useEffect(() => {
    getUsers().then(users => {
      setUsers(users)
    })
  }, [users, posts]);

  return (
    <div className="App">
      {displayUsers(users)}
    </div>
  );
}

type User = {
  id: number;
  name: string;
  email: string;
  password: string;
  age: number;
  createdAt: string;
}

type Post = {
  id: number;
  userEmail: string;
  text: string;
  createdAt: string;
}