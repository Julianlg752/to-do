import { useState } from 'react';
import styles from '../styles/LoginPage.module.css'
import { useRouter } from 'next/router';

const LoginPage = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const router = useRouter();
    const endpoint = process.env.API_URL;
    const handleSubmit = async (e) => {
        e.preventDefault();
        // Realizar solicitud POST al backend para iniciar sesión y obtener el token JWT
        try {
            const response = await fetch(endpoint+'/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, password }),
            });
            const data = await response.json();

            if (response.ok) {
                // Guardar el token JWT en el almacenamiento local del navegador (localStorage o sessionStorage)
                localStorage.setItem('token', data.token);
                // Redireccionar a la página de inicio
                router.push("/dashboard");
            } else {
                setError(data.error);
            }
        } catch (error) {
            console.log(error)
            setError("Internal Server Error");
        }
    };
    
    return (
        <div className={styles.loginPage}>
            <h1 className={styles.title}>Iniciar sesión</h1>
            <form onSubmit={handleSubmit} className={styles.loginForm}>
                <input
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    placeholder="Usuario"
                    className={styles.input}
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Contraseña"
                    className={styles.input}
                />
                <button type="submit" className={styles.loginButton}>
                    Iniciar sesión
                </button>
                {error && <p className={styles.error}>{error}</p>}
            </form>
        </div>
    );
};


export default LoginPage;
