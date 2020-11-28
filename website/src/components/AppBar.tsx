import React from 'react';
import Container from "@material-ui/core/Container";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import LockOpenIcon from "@material-ui/icons/LockOpen";
import Typography from "@material-ui/core/Typography";
import {Link} from "@material-ui/core";
import Button from "@material-ui/core/Button";
import makeStyles from "@material-ui/core/styles/makeStyles";
import {AppBar as MaterialAppBar} from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
    header: {
        flex: 0,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
}));

function AppBar() {
    const classes = useStyles();
    return (
        <React.Fragment>
            <Container maxWidth={false} disableGutters className={classes.header}>
                <MaterialAppBar position="static">
                    <Container maxWidth="sm" disableGutters>
                        <Toolbar>
                            <IconButton edge="start" className={classes.menuButton} color="inherit"
                                        aria-label="menu">
                                <LockOpenIcon/>
                            </IconButton>
                            <Typography variant="h6" className={classes.title}>
                                Privacy
                            </Typography>
                            <Link href="https://github.com/kallydev/privacy" underline="none" target="_blank"
                                  component={Button}
                                  color="inherit">GitHub</Link>
                        </Toolbar>
                    </Container>
                </MaterialAppBar>
            </Container>
        </React.Fragment>
    );
}

export default AppBar;
