import async from './fetch';

import apiPost from './api-post';

const baseUrl = '/user';

export default {
    login(username,password){
        return async(
            `${baseUrl}/login`,{
                username:username,
                password:password
            },
            'post'
        );
    },
    signOut(){
        return async(`${baseUrl}/signOut`);
    },
    ...apiPost,
};