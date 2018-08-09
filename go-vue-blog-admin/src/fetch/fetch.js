import axios from 'axios';

// 需要使用代理来解决跨域问题
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
axios.defaults.timeout = 20000;
axios.defaults.baseURL = 'http://127.0.0.1:8081/v1'
// Add a request interceptor
axios.interceptors.request.use(
    config => {
        const token = localStorage.getItem('BLOG_TOKEN');
        if (token){
            // Louyan是JWT的认证头部信息
            config.headers.common['Authorization'] = 'Louyan ' + token;
        }
        return config;
    }
);

// Add a response interceptor

axios.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        return Promise.reject(error);
    }
);


export default async(url = '',params = {},method = 'get',isUpload = false) => {
    method = method.toLowerCase();
    if (method === 'get'){
        let paramArr = [];
        for (let [key,value] of Object.entries(params)){
            paramArr.push(key + '=' + value);
        }
        if (paramArr.length > 0){
            url += '?' + paramArr.join('&').replace(/#/g,'%23');
        }

        return new Promise((resolve,reject)=>{
            axios.get(url)
                .then(
                    response => {
                        resolve(response.data);
                    },
                    err => {
                        reject(err);
                    }
                )
                .catch(errror => {
                    reject(errror);
                });
        });
    }else if (method ==='post') {
        let config = {};
        if (isUpload){
            config = {
                headers:{
                    'Content-Type': 'multipart/form-data'
                }
            };
        }

        return new Promise((resolve,reject)=>{
            axios.post(url,params,config)
                .then(
                    response => {
                        resolve(response.data);
                    },
                    err => {
                        reject(err);
                    }
                )
                .catch(errror => {
                    reject(errror);
                });
        });
    }else if (method === 'put'){
        return new Promise((resolve,reject)=>{
            axios.put(url,params)
                .then(
                    response => {
                        resolve(response.data);
                    },
                    err => {
                        reject(err);
                    }
                )
                .catch(errror => {
                    reject(errror);
                });
        });
    } else if (method === 'delete'){
        return new Promise((resolve,reject)=>{
            axios.delete(url)
                .then(
                    response => {
                        resolve(response.data);
                    },
                    err => {
                        reject(err);
                    }
                )
                .catch(errror => {
                    reject(errror);
                });
        });
    } else {
        let error='传递的参数错误';
        return Promise.reject(error);
    }

};














