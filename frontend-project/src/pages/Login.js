import React from 'react';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import './Login.css'
import user_icon from '../Assets/person.png'
import email_icon from '../Assets/email.png'
import password_icon from '../Assets/password.png'


const Login = () => {
  const navigate = useNavigate();

  const formik = useFormik({
    initialValues: {
      username: '',
      password: '',
    },
    validationSchema: Yup.object({
      username: Yup.string().required('Required'),
      password: Yup.string().required('Required'),
    }),
    onSubmit: async (values) => {
      try {
        const response = await axios.post('http://localhost:8080/login', values);
        localStorage.setItem('token', response.data.token);
        navigate('/authenticated');
      } catch (error) {
        alert('Login failed');
      }
    },
  });

  return (
    <div className='container'>
    <form onSubmit={formik.handleSubmit}>
      
        <div className="header">
            <div className="text">Login</div>
            <div className="underline"></div>
        </div>

        <div className="input">
        <img src={user_icon} alt="" />
        <input
          id="username"
          name="username"
          type="text" 
          placeholder = "Username"
          onChange={formik.handleChange}
          value={formik.values.username}
        />
        {formik.touched.username && formik.errors.username ? (
          <div>{formik.errors.username}</div>
        ) : null}
        </div>


      <div className="input">
        <img src={password_icon} alt="" />
        <input
          id="password"
          name="password"
          type="password"
          placeholder = "Password"
          onChange={formik.handleChange}
          value={formik.values.password}
        />
        {formik.touched.password && formik.errors.password ? (
          <div>{formik.errors.password}</div>
        ) : null}
      </div>


        <div className="submit-container">
            <button type="submit">Login</button>
        </div>
    
    </form>
    </div>
  );
};

export default Login;
