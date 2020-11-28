import Box from '@material-ui/core/Box';
import Paper from '@material-ui/core/Paper';
import React from 'react';
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import {List, ListItem, ListItemIcon, ListItemText, SvgIcon} from "@material-ui/core";
import PhoneIcon from "@material-ui/icons/Phone";
import {QueryResult, State} from '../App';
import FaceIcon from '@material-ui/icons/Face';
import DoneIcon from '@material-ui/icons/Done';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import EmailIcon from '@material-ui/icons/Email';
import BusinessIcon from '@material-ui/icons/Business';
import VpnKeyIcon from '@material-ui/icons/VpnKey';

interface ResultProps {
    state: State;
    closeResult: () => void;
    result: QueryResult;
}

function Result(props: ResultProps) {
    return (
        <React.Fragment>
            {
                props.state === State.Succeeded &&
                <Paper>
                    <Box pt={2} px={2} pb={1}>
                        <Box display="flex" flexDirection="row" justifyContent="center">
                            <Box flexGrow={1}>
                                <Typography component="h5" variant="h5">检测结果</Typography>
                            </Box>
                            <IconButton size="small">
                                <CloseIcon onClick={() => {
                                    props.closeResult();
                                }}/>
                            </IconButton>
                        </Box>
                        <Box mt={1}>
                            <List disablePadding component="nav" aria-label="main mailbox folders">
                                {
                                    (
                                        props.result.names.length === 0 &&
                                        props.result.nicknames.length === 0 &&
                                        props.result.phone_numbers.length === 0 &&
                                        props.result.id_numbers.length === 0 &&
                                        props.result.qq_numbers.length === 0 &&
                                        props.result.passwords.length === 0 &&
                                        props.result.emails.length === 0 &&
                                        props.result.addresses.length === 0
                                    ) && <ListItem disableGutters>
                                        <ListItem>
                                            <ListItemIcon>
                                                <DoneIcon/>
                                            </ListItemIcon>
                                            <ListItemText primary="未查询到任何结果"/>
                                        </ListItem>
                                    </ListItem>
                                }
                                {
                                    props.result.names.map((name, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <FaceIcon/>
                                                </ListItemIcon>
                                                <ListItemText primary={name}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.nicknames.map((nickname, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <AccountCircleIcon/>
                                                </ListItemIcon>
                                                <ListItemText primary={nickname}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.phone_numbers.map((phoneNumber, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <PhoneIcon/>
                                                </ListItemIcon>
                                                <ListItemText primary={phoneNumber}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.id_numbers.map((idNumber, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <SvgIcon>
                                                        <path fill="currentColor"
                                                              d="M2,3H22C23.05,3 24,3.95 24,5V19C24,20.05 23.05,21 22,21H2C0.95,21 0,20.05 0,19V5C0,3.95 0.95,3 2,3M14,6V7H22V6H14M14,8V9H21.5L22,9V8H14M14,10V11H21V10H14M8,13.91C6,13.91 2,15 2,17V18H14V17C14,15 10,13.91 8,13.91M8,6A3,3 0 0,0 5,9A3,3 0 0,0 8,12A3,3 0 0,0 11,9A3,3 0 0,0 8,6Z"/>
                                                    </SvgIcon>
                                                </ListItemIcon>
                                                <ListItemText primary={idNumber}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.qq_numbers.map((qqNumber, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <SvgIcon>
                                                        <path fill="currentColor"
                                                              d="M3.18,13.54C3.76,12.16 4.57,11.14 5.17,10.92C5.16,10.12 5.31,9.62 5.56,9.22C5.56,9.19 5.5,8.86 5.72,8.45C5.87,4.85 8.21,2 12,2C15.79,2 18.13,4.85 18.28,8.45C18.5,8.86 18.44,9.19 18.44,9.22C18.69,9.62 18.84,10.12 18.83,10.92C19.43,11.14 20.24,12.16 20.82,13.55C21.57,15.31 21.69,17 21.09,17.3C20.68,17.5 20.03,17 19.42,16.12C19.18,17.1 18.58,18 17.73,18.71C18.63,19.04 19.21,19.58 19.21,20.19C19.21,21.19 17.63,22 15.69,22C13.93,22 12.5,21.34 12.21,20.5H11.79C11.5,21.34 10.07,22 8.31,22C6.37,22 4.79,21.19 4.79,20.19C4.79,19.58 5.37,19.04 6.27,18.71C5.42,18 4.82,17.1 4.58,16.12C3.97,17 3.32,17.5 2.91,17.3C2.31,17 2.43,15.31 3.18,13.54Z"/>
                                                    </SvgIcon>
                                                </ListItemIcon>
                                                <ListItemText primary={qqNumber}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.passwords.map((qqNumber, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <VpnKeyIcon/>
                                                </ListItemIcon>
                                                <ListItemText primaryTypographyProps={{style: {wordWrap: "break-word"}}}
                                                              primary={qqNumber}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.emails.map((email, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <EmailIcon/>
                                                </ListItemIcon>
                                                <ListItemText primary={email}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                                {
                                    props.result.addresses.map((address, index) =>
                                        <ListItem disableGutters key={index}>
                                            <ListItem>
                                                <ListItemIcon>
                                                    <BusinessIcon/>
                                                </ListItemIcon>
                                                <ListItemText primary={address}/>
                                            </ListItem>
                                        </ListItem>
                                    )
                                }
                            </List>
                        </Box>
                    </Box>
                </Paper>
            }
        </React.Fragment>
    );
}

export default Result;
