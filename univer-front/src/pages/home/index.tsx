import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

interface User {
    user_id: number;
    username: string;
    exp: number;
}

const HomePage = () => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUser = async () => {
            const token = localStorage.getItem('token');
            if (!token) {
                navigate('/login');
                return;
            }

            try {
                const response = await axios.get('/api/users/me', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                if (response.data.status === 'success') {
                    setUser(response.data.data);
                }
            } catch (error) {
                console.error('Failed to fetch user:', error);
                localStorage.removeItem('token');
                navigate('/login');
            } finally {
                setLoading(false);
            }
        };

        fetchUser();
    }, [navigate]);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    if (loading) {
        return <div className="">Loading...</div>;
    }

    return (
        <div className="">
            
                {user ? (
                        <>
                            <p className="text-gray-900 text-lg font-bold">{user.username}</p>
                        

                        
                            <p className="text-gray-900 text-lg font-bold">{user.user_id}</p>
                        

                        <button
                            onClick={handleLogout}
                            
                        >
                            Logout
                        </button>
                        </>
                ) : (
                    <div className="">Failed to load user data</div>
                )}
            
        </div>
    );
};

export default HomePage;
