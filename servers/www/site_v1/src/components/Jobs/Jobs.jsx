import React, {useEffect} from "react";
import * as api from "../../api/api";
import { useHistory } from "react-router-dom";
import JobNew from "./JobNew";
import JobList from "./JobList";

export default function Jobs() {

    const history = useHistory()
    const [jobs, setJobs] = React.useState([])

    function getJobs() {
        api.getJobs().then(response => {
            setJobs(response.data)
        })
    }

    useEffect(() => {
        getJobs()
    }, [])

    function onClickJob(e) {
        //console.log("Click: " + e.target.id)
    }

    function onClickJobClose(e) {
        //console.log("Close: " + e.target.id)
        api.closeJob(e.target.id).then(response => {
            getJobs()
        })
    }

    function onClickJobOpen(e) {
        //console.log("Open: " + e.target.id)
        history.push('/job/' + e.target.id)
    }

    function createdJob(job_id) {
        history.push('/job/' + job_id)
    }

    return (
        <div>
            <h2>Job</h2>
            <hr/>

            <JobNew created={createdJob}/>

            <JobList jobs={jobs}
                     onClickJob={onClickJob}
                     onClickJobCloce={onClickJobClose}
                     onClickJobOpen={onClickJobOpen}
            />
        </div>
    )
}
