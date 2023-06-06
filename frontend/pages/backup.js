import { useState, useEffect } from 'react';
import axios from 'axios';
import styles from '../styles/DasboardPage.module.css';
import { useRouter } from 'next/router';

const DashboardPage = () => {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState('');
  const [newTaskDescription, setNewTaskDescription] = useState('');
  const endpoint = process.env.API_URL;
  const ISSERVER = typeof window === "undefined";
  
  let token = "";
  const router = useRouter();

  if (!ISSERVER) {
    token = localStorage.getItem('token');
  }
  useEffect(() => {
    if (!token) {
      router.push("/");
    }
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {

      const response = await axios.get(endpoint + '/api/tasks', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setTasks(response.data);
    } catch (error) {
      console.log(error);
    }
  };

  const handleCheckboxChange = async (taskId, completed) => {
    try {
      await axios.put(endpoint + `/api/tasks/${taskId}`, {
        title: newTask,
        description: newTaskDescription,
        status: completed,
      },
        {
          headers: { Authorization: `Bearer ${token}` },
        });

      setTasks((prevTasks) =>
        prevTasks.map((task) => {
          if (task.id === taskId) {
            return { ...task, completed };
          }
          return task;
        })
      );
    } catch (error) {
      console.log(error);
    }
  };

  const handleDeleteTask = async (taskId) => {
    try {
      await axios.delete(endpoint + `/api/tasks/${taskId}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setTasks((prevTasks) => prevTasks.filter((task) => task.id !== taskId));
    } catch (error) {
      console.log(error);
    }
  };

  const handleUpdateTask = async (taskId, newName, newDescription, status) => {
    try {
      await axios.put(endpoint + `/api/tasks/${taskId}`,
        {
          title: newName,
          description: newDescription,
          status: status
        }, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setTasks((prevTasks) =>
        prevTasks.map((task) => {
          if (task.id === taskId) {
            return { ...task, name: newName, description: newDescription };
          }
          return task;
        })
      );
    } catch (error) {
      console.log(error);
    }
  };

  const handleAddTask = async () => {
    try {
      const response = await axios.post(endpoint + '/api/tasks', {
        title: newTask,
        description: newTaskDescription,
        status: true,
      }, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setTasks((prevTasks) => [...prevTasks, response.data]);
      setNewTask('');
      setNewTaskDescription('');
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className={styles.dashboard}>
      <h1 className={styles.title}>Dashboard</h1>
      <div className={styles.taskList}>
        <h2 className={styles.subtitle}>Tareas</h2>
        <ul>
          {tasks.map((task) => (
            <li key={task.id} className={styles.taskItem}>
              <input
                type="checkbox"
                checked={task.completed}
                onChange={(e) => handleCheckboxChange(task.id, e.target.checked)}
              />
              <div>
                <h3 className={styles.taskTitle}>{task.title}</h3>
                <p className={styles.taskDescription}>{task.description}</p>
              </div>
              <div className={styles.taskActions}>
                <button
                  onClick={() => handleDeleteTask(task.id)}
                  className={styles.deleteButton}
                >
                  Eliminar
                </button>
                <button
                  onClick={() =>
                    handleUpdateTask(task.id, `${task.name} (actualizado)`, task.description)
                  }
                  className={styles.updateButton}
                >
                  Actualizar
                </button>
              </div>
            </li>
          ))}
        </ul>
      </div>
      <div className={styles.newTask}>
        <h2 className={styles.subtitle}>Agregar Nueva Tarea</h2>
        <input
          type="text"
          placeholder="Título"
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          className={styles.newTaskInput}
        />
        <textarea
          placeholder="Descripción"
          value={newTaskDescription}
          onChange={(e) => setNewTaskDescription(e.target.value)}
          className={styles.newTaskTextArea}
        ></textarea>
        <button onClick={handleAddTask} className={styles.addButton}>
          Agregar Tarea
        </button>
      </div>
    </div>
  );
};


export default DashboardPage;
