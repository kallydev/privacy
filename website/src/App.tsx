import React from 'react';
import 'fontsource-roboto';
import './App.css';
import Typography from "@material-ui/core/Typography";
import Button from '@material-ui/core/Button';
import makeStyles from '@material-ui/core/styles/makeStyles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import TextField from "@material-ui/core/TextField";
import FormControl from '@material-ui/core/FormControl';
import Result from "./components/ResultList";
import AppBar from "./components/AppBar";
import Footer from "./components/Footer";
import Alert from '@material-ui/lab/Alert';

const useStyles = makeStyles((theme) => ({
    main: {
        flex: 1,
        padding: theme.spacing(2),
    },
    heroContent: {
        padding: theme.spacing(4, 2),
    },
}));

interface Response {
    code: number;
    message: string;
    data: any;
}

export interface QueryResult {
    names: string[];
    nicknames: string[];
    phone_numbers: string[];
    id_numbers: string[];
    qq_numbers: string[];
    passwords: string[];
    emails: string[];
    addresses: string[];
}

export enum State {
    Normal,
    Succeeded,
    Failed,
}

function App() {
    const classes = useStyles();

    const [value, setValue] = React.useState('');
    const [state, setState] = React.useState(State.Normal);
    const [error, setError] = React.useState('');
    const [result, setResult] = React.useState({
        names: [],
        nicknames: [],
        phone_numbers: [],
        id_numbers: [],
        qq_numbers: [],
        passwords: [],
        emails: [],
        addresses: [],
    } as QueryResult);

    const closeResult = () => {
        setState(State.Normal);
    }

    const query = async () => {
        setState(State.Normal);
        if (value.length === 0) {
            return;
        }
        let url = new URL(document.baseURI + 'api/query');
        url.search = new URLSearchParams({'value': value}).toString();
        fetch(url.toString(), {})
            .then(response => response.json())
            .then((response: Response) => {
                if (response.code !== 0) {
                    setState(State.Failed);
                    setError(response.message)
                    return;
                }
                setResult(response.data as QueryResult);
                setState(State.Succeeded);
            })
            .catch((error) => {
                setState(State.Failed);
                setError(error)
                console.log('error: ' + error);
            });
    }

    return (
        <React.Fragment>
            <CssBaseline/>
            <AppBar/>
            <Container maxWidth="sm" component="main" className={classes.main}>
                {
                    state === State.Failed &&
                    <Alert variant="filled" severity="error">
                        {error}
                    </Alert>
                }
                <Box mt={2}>
                    <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                        Privacy
                    </Typography>
                    <Typography component="p" variant="h5" align="center" color="textSecondary">
                        「 is dead, get over it. 」
                    </Typography>
                </Box>
                <Box mt={4}>
                    <FormControl fullWidth component="fieldset">
                        <TextField
                            fullWidth
                            id="filled-basic"
                            label="QQ / 手机号 / 身份证号 / 邮箱"
                            variant="outlined"
                            type="number"
                            onChange={event => {
                                setValue(event.target.value)
                            }}
                            color="secondary"/>
                        <Box mt={2}>
                            <Button fullWidth size="large" variant="contained" color="secondary" onClick={query}>
                                检测隐私状态
                            </Button>
                        </Box>
                    </FormControl>
                </Box>
                {
                    result !== null &&
                    <Box mt={4}>
                        <Result state={state} closeResult={closeResult} result={result}/>
                    </Box>
                }
            </Container>
            <Footer/>
        </React.Fragment>
    );
}

export default App;
