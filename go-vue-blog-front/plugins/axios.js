import * as axios from 'axios';

export default () => {
  axios.defaults.baseURL = 'http://localhost:8081/v1';
  // axios.defaults.timeout = 5000;
  axios.defaults.headers.post['Content-Type'] =
    'application/x-www-form-urlencoded;charset=UTF-8';
  axios.interceptors.request.use(
    config => {
      return config;
    },
    err => {
      return Promise.reject(err);
    }
  );

  axios.interceptors.response.use(
    response => {
      return response;
    },
    err => {
      return Promise.reject(err);
    }
  );
};
