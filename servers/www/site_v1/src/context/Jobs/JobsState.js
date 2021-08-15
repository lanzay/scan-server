import React, {useReducer} from 'react'
import * as api from '../../api/api';

import {JobsContext} from './JobsContext'
import {JobsReducer} from './JobsReducer'

import {GET_JOBS} from "../types";

export const JobsState = ({children}) => {

    const initialState = {
        jobs: [],
    }

    const [state, dispatch] = useReducer(JobsReducer, initialState);

    const getJobs = (days) => {
        const functionRequest = () => {
            return api.getJobs(days)
        };
        const responseHandlingFunction = (json) => {
            if (!json.error) {
                dispatch({type: GET_JOBS, payload: {json}})
            } else {
                dispatch({type: GET_JOBS, payload: {json}})
            }
        }
        const exceptionHandlingFunction = (err) => {
            dispatch({type: GET_JOBS, err: {err}})
        };
        api.executorRequests(functionRequest, responseHandlingFunction, exceptionHandlingFunction, dispatch);
    };


    return (
        <JobsContext.Provider value={{
            jobs: state.jobs,
            getJobs,
        }}>{children}
        </JobsContext.Provider>
    )
}
