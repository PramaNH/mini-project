import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useFormik } from 'formik';
import * as Yup from 'yup';

const Authenticated = () => {
  const navigate = useNavigate();
  const [data, setData] = useState([]); // Holds fetched data
  const [loading, setLoading] = useState(true); // Loading state for API call
  const [error, setError] = useState(null); // Error state for API call

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/data', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        setData(response.data); // Set the fetched data
      } catch (err) {
        setError('Failed to fetch data. Redirecting to login...');
        setTimeout(() => navigate('/login'), 2000); // Redirect to login after 2 seconds
      } finally {
        setLoading(false); // Stop loading regardless of success or error
      }
    };

    fetchData();
  }, [navigate]);

  const formik = useFormik({
    initialValues: {
      name: '',
      email: '',
    },
    validationSchema: Yup.object({
      name: Yup.string().required('Name is required'),
      email: Yup.string().email('Invalid email address').required('Email is required'),
    }),
    onSubmit: async (values) => {
      try {
        await axios.post('http://localhost:8080/data', values, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        alert('Data submitted successfully');
        // Refresh the data after submission
        const response = await axios.get('http://localhost:8080/data', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        setData(response.data);
      } catch (err) {
        alert('Failed to submit data');
      }
    },
  });

  return (
    <div>
      <h1>Authenticated Page</h1>


      {loading ? (
        <p>Loading...</p>
      ) : error ? (
        <p>{error}</p>
      ) : (
        <>
          {/* Data List */}
          <ul>
            {data.length > 0 ? (
              data.map((item) => (
                <li key={item.ID}>
                  {item.Name} - {item.Email}
                </li>
              ))
            ) : (
              <p>No data available</p>
            )}
          </ul>

          {/* Data Submission Form */}
          <form onSubmit={formik.handleSubmit}>
            <div>
              <label htmlFor="name">Name</label>
              <input
                id="name"
                name="name"
                type="text"
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                value={formik.values.name}
              />
              {formik.touched.name && formik.errors.name && (
                <div style={{ color: 'red' }}>{formik.errors.name}</div>
              )}
            </div>
            <div>
              <label htmlFor="email">Email</label>
              <input
                id="email"
                name="email"
                type="email"
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                value={formik.values.email}
              />
              {formik.touched.email && formik.errors.email && (
                <div style={{ color: 'red' }}>{formik.errors.email}</div>
              )}
            </div>
            <button type="submit">Submit</button>
          </form>
        </>
      )}
    </div>
  );
};

export default Authenticated;
