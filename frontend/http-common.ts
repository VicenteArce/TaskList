import axios from 'axios';

const todoListServer = process.env.TODO_LIST_SERVER || 'localhost:8080';

export default axios.create({
  baseURL: `http://${todoListServer}`,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: false 
});