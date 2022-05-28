import React from 'react';
import axios from 'axios';

export default function App() {

  const [users, setUsers] = React.useState<User | null | undefined>(undefined);
  // const [posts, setPosts] = React.useState<Post[] | null>([]);

  let displayUsers = (user: User | null | undefined) => {
    if (user!==null) {
      return <h1>Hello user: {user?.name}</h1>
    }
    // if (users?.length! > 1 && users!==null) {
    //   for (const user of users) {
    //     return <h1>Hello user: {user.name}</h1>
    //   }
    // }
  }
  
  async function getUsers(): Promise<User | null | undefined> {
    const { data, status } = await axios.get<User | null | undefined>('http://localhost:8080/users/cristian@mina.com')
    if (status === 200) {
      return data
    } else {
      console.log('error')
      return null
    }
  }

  // async function getPosts(user: User) {
  //   const { data, status } = await axios.get<Post[] | null>('http://localhost:8080/users')
  //   if (status === 200) {
  //     console.log(data)
  //     return data
  //   } else {
  //     console.log('error')
  //     return null
  //   }
  // }

  React.useEffect(() => {
      getUsers().then(user => {
        setUsers(user);
      })
      console.table(users);
  }, [users]);

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

// type Post = {
//   id: number;
//   userEmail: string;
//   text: string;
//   createdAt: string;
// }