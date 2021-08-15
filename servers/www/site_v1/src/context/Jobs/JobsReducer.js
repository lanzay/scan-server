import {GET_JOBS} from '../types'

export const JobsReducer = (state, action) => {

    switch (action.type) {

        case GET_JOBS:
            return {
                ...state,
                ...action.payload,
            };
        default:
            return state
    }
}
