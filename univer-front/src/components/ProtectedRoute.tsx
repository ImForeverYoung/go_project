import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import axios from 'axios';

interface ProtectedRouteProps {
    children: React.ReactNode;
}

const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
    const token = localStorage.getItem('token');
    const [isValid, setIsValid] = useState<boolean | null>(null);

    useEffect(() => {
        const verifyToken = async () => {
            if (!token) {
                setIsValid(false);
                return;
            }

            try {
                await axios.get('/api/users/me', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                setIsValid(true);
            } catch (error) {
                console.error('Token verification failed:', error);
                localStorage.removeItem('token');
                setIsValid(false);
            }
        };

        verifyToken();
    }, [token]);

    if (isValid === null) {
        
        return <div className="min-h-screen flex items-center justify-center">Loading...</div>;
    }

    if (!isValid) {
        return <Navigate to="/login" replace />;
    }

    return <>{children}</>;
};

export default ProtectedRoute;
