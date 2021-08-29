import axios from 'axios'

//const API_URL = 'https://app.swaggerhub.com/apis/elisru/ret-shop/1.0.0'
const API_URL = 'http://localhost:3030/api/v1.0/'

export const executorRequests = (functionRequest, responseHandlingFunction, exceptionHandlingFunction, dispatch, repeat = false) => {

    axios.defaults.headers.common = {
        'authorization': `bearer ${'112233'}`,
    }

    functionRequest()
        .then(response => responseHandlingFunction(response.data))
        .catch((error) => {

            console.log('error', {error});

            if (!error.response && !error.request) {
                exceptionHandlingFunction("что то полшло не так..." + error.message);
            } else if (!error.response && error.request) {
                exceptionHandlingFunction("Проблема соединения")
            } else if (error.request.status === 401) {

                if (repeat) {
                    //dispatch(logOut())
                } else {
                    // refreshToken()
                    //     .then(response => {
                    //         setAccessToken(response.data.accessToken)
                    //         return executorRequests(functionRequest, responseHandlingFunction, exceptionHandlingFunction, dispatch, true);
                    //     })
                    //     .catch((error) => {
                    //         //dispatch(logOut())
                    //     })
                }

            } else {
                exceptionHandlingFunction(error.response.data)
            }
        })
}

export const getJobs = () => {
    return axios.get(`${API_URL}jobs`);
}

export const newJob = (job_name, comment) => {
    return axios.post(`${API_URL}job/new`, {job_name, comment});
}

export const getJob = (job_id) => {
    return axios.get(`${API_URL}job/${job_id}`);
}

export const scanBarcode = (job_id, barcode, count = 1) => {
    return axios.put(`${API_URL}job/${job_id}/scan/${barcode}/${count}`);
}

export const closeJob = (job_id) => {
    return axios.get(`${API_URL}job/${job_id}/close`);
}

//========================
export const getAccessToken = () => {
    return localStorage.getItem('accessToken')
}

export const setAccessToken = (accessToken) => {
    console.log('accessToken', accessToken);
    localStorage.setItem('accessToken', accessToken);
}


