import React, {useEffect} from "react";
import * as api from "../../api/api";
import JobList from "./JobList";

export default function Jobs() {

    const [jobs, setJobs] = React.useState([])
    const [job, setJob] = React.useState({
        job_name: "",
        comment: "",
    })

    useEffect(() => {
        api.getJobs(1).then(response => {
            setJobs(response.data)
        })
    }, [])

    function getJobs() {
        api.getJobs(1).then(response => {
            setJobs(response.data)
        })
    }

    function newJob() {
        api.newJob('qqq', 'comment 1').then(response => {
            setJobs([response.data])
        })
    }

    function onClickJob(e) {
        // TODO global state
        console.log("Click: " + e.target.id)
    }

    function onClickJobCloce(e) {
        // TODO global state
        console.log("Close: " + e.target.id)
    }
    function onClickJobOpen(e) {
        // TODO global state
        console.log("Open: " + e.target.id)
    }

    return (
        <div>
            <h2>Jobs</h2>
<hr/>

            <h4 >Новое задание</h4>
            <div className="mb-3">
                <label for="job_name" className="form-label">Название:</label>
                <input id="job_name" type="text" className="form-control"
                       value={job.job_name}
                       onInput={(v) => setJob({...job, job_name: v.target.value})}/>
                <label for="job_comment" className="form-label">Комментарий:</label>
                <textarea id="job_comment" type="text" className="form-control"
                       value={job.comment}
                       onInput={(v) => setJob({...job, comment: v.target.value})}/>
                <div className="d-grid gap-2 mt-3">
                <button onClick={newJob} type="button" className="btn btn-primary">Создать</button>
                </div>
            </div>

            <h2>Job list (open)</h2>
            <JobList jobs={jobs}
                     onClickJob={onClickJob}
                     onClickJobCloce={onClickJobCloce}
                     onClickJobOpen={onClickJobOpen}
            />
        </div>
    )
}
