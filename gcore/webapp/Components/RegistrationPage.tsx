import * as React from "react";

import { NURL, HTTPClient } from "tsutils";

import lightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import AppBar from 'material-ui/AppBar';
import TextField from 'material-ui/TextField';
import Paper from 'material-ui/Paper';
import RaisedButton from 'material-ui/RaisedButton';
import CircularProgress from 'material-ui/CircularProgress';

import { Theme } from "./Theme";
export { RegistrationPage };

let paperStyle = {
    display: "block",
    margin: "auto",
    marginTop: "90px",
    width: "320px",
    padding: "20px",
    textAlign: "center"
}

let formStyle = {
    color: "#333",
    display: "block",
    margin: "auto",
    marginTop: "20px",
    marginBottom: "20px",
}

let baseURL = new NURL(window.location.href);

interface RPState {
    captchaID: string;
    username: string;
    password: string;
    usernameError?: string;
    passwordError?: string;
    captchaError?: string;
    captchaSolution: string;
    submitting: boolean;
    done: boolean;
}

interface RPProps { }

class RegistrationPage extends React.Component<RPProps, RPState> {
    constructor(props) {
        super(props);
        this.state = { done: false, submitting: false, captchaSolution: "", captchaID: "", username: "", password: "", usernameError: undefined };
        this.setCaptcha();
    }

    setCaptcha() {
        let cp = baseURL.subPath("/v1/newCaptcha").serialize();
        new HTTPClient().getJSON(cp).then((resp) => {
            let id = encodeURIComponent(resp["captchaID"])
            this.setState({
                captchaID: id
            });
        });
    }

    handleCaptcha(cap): void {
        let ct = cap.target.value;
        this.setState({
            captchaError: undefined,
            captchaSolution: ct
        });
    }

    handlePass(cap): void {
        let ct = cap.target.value;
        this.setState({
            password: ct,
            passwordError: undefined
        });
    }

    handleUsername(usr): void {
        let username = usr.target.value;
        this.setState({
            username: username,
            usernameError: ""
        });

        if (username === "") {
            return;
        }

        new HTTPClient().getJSON(baseURL.subPath(`/v1/userExists/${username}`).serialize()).then((resp) => {
            let err;
            if (resp["exists"]) {
                err = "Username in use.";
            }

            this.setState({
                usernameError: err
            });
        });
    }

    submit() {
        if (this.state.password === "") {
            this.setState({
                passwordError: "You must enter a password."
            });
            return;
        }
        this.setState({
            submitting: true
        });
        new HTTPClient().getJSON(baseURL.subPath(`/v1/register/${this.state.username}/${this.state.password}/${this.state.captchaID}/${this.state.captchaSolution}`).serialize())
            .then((resp) => {
                if (resp["error"]) {
                    this.setState({
                        submitting: false
                    });

                    if (/captcha/.test(resp["error"])) {
                        this.setState({
                            captchaError: resp["error"]
                        });

                        this.setCaptcha();
                        return;
                    }

                    this.setState({
                        usernameError: resp["error"]
                    });
                    return;
                }

                this.setState({
                    done: true
                });
            });
    }

    render() {
        if (this.state.submitting) {
            return (
                <MuiThemeProvider muiTheme={getMuiTheme(Theme)}>
                    <div>
                        <AppBar title="Gophercraft" />
                        <Paper style={paperStyle} zDepth={2}>
                            <h2 style={{ fontFamily: "Roboto", fontWeight: 400 }}>Submitting</h2>
                            <Spinner done={this.state.done}></Spinner>
                        </Paper>
                    </div>
                </MuiThemeProvider>
            )
        }

        return (
            <MuiThemeProvider muiTheme={getMuiTheme(Theme)}>
                <div>
                    <AppBar title="Gophercraft" />

                    <Paper style={paperStyle} zDepth={2}>
                        <h2 style={{ fontFamily: "Roboto", fontWeight: 400 }}>Registration</h2>

                        <TextField onChange={this.handleUsername.bind(this)} style={formStyle} errorText={this.state.usernameError} hintText="username" value={this.state.username}></TextField>
                        <TextField onChange={this.handlePass.bind(this)} style={formStyle} errorText={this.state.passwordError} hintText="password" type="password"></TextField>

                        <img src={baseURL.subPath(`/v1/captcha/${this.state.captchaID}.png`).serialize()} />

                        <TextField onChange={this.handleCaptcha.bind(this)} style={formStyle} errorText={this.state.captchaError} hintText="solve captcha"></TextField>
                        <br />

                        <RaisedButton onClick={this.submit.bind(this)} style={{ width: "230px" }} label="Submit" primary={true} />
                    </Paper>
                </div>
            </MuiThemeProvider>
        )
    }
}

interface SpinnerProps {
    done: boolean;
}

let successStyle = {
    textAlign: "center",
    fontFamily: "Roboto"
}

class Spinner extends React.Component<SpinnerProps, any> {
    constructor(props) {
        super(props);
    }

    render() {
        if (this.props.done) {
            return (
                <div>
                    <p style={successStyle}>Success!</p>
                    <p style={successStyle}>✓</p>
                </div>
            )
        }

        return (
            <CircularProgress size={60} thickness={7} />
        )
    }
}