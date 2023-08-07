import { useEffect, useState } from "react";
import {  useSearchParams} from "react-router-dom";
import "./App.css";
import Backdrop from '@mui/material/Backdrop';
import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import axios from "axios";

function App() {
  const [walletAddress, setWalletAddress] = useState("");
  const [status, setStatus] = useState();
  const [boolMM, setBoolMM] = useState(true)
  const [searchParams, setSearchParams] = useSearchParams();
   const [open, setOpen] = useState(false);
  const handleClose = () => {
    setOpen(false);
  };
  const handleOpen = () => {
    setOpen(true);
  };
  useEffect(() => {
    getCurrentWalletConnected();
    addWalletListener();
  }, []);

  const connectWallet = async () => {
    if (typeof window != "undefined" && typeof window.ethereum != "undefined") {
      try {
        const accounts = await window.ethereum.request({
          method: "eth_requestAccounts",
        });
        setWalletAddress(accounts[0]);
      } catch (err) {
        console.error(err.message);
      }
    }
  };

  const getCurrentWalletConnected = async () => {
    if (typeof window != "undefined" && typeof window.ethereum != "undefined") {
      try {
        const accounts = await window.ethereum.request({
          method: "eth_accounts",
        });
        if (accounts.length > 0) {
          setWalletAddress(accounts[0]);
          console.log(accounts[0]);
        } else {
          console.log("Connect to MetaMask using the Connect button");
        }
      } catch (err) {
        console.error(err.message);
      }
    } else {
      /* MetaMask is not installed */
      console.log("Please install MetaMask");
    }
  };

  const addWalletListener = async () => {
    if (typeof window != "undefined" && typeof window.ethereum != "undefined") {
      window.ethereum.on("accountsChanged", (accounts) => {
        setWalletAddress(accounts[0]);
        console.log(accounts[0]);
      });
    } else {
      /* MetaMask is not installed */
      setBoolMM(false);
    }
  };

  const signMessage = async () =>{
    const signedData = await window.ethereum.request({
      method: "personal_sign",
      params: [
        window.atob(searchParams.get("sign")),
        walletAddress,
      ],
    });
    

    console.log(signedData)
    axios.post("http://localhost:8081/connect", {
      message: searchParams.get("sign"),
      wallet: walletAddress,
      signed_message: signedData,
    })

    .then(response => {
    console.log(response.status);
    setStatus(response.status);
    setOpen(true);
    })

    .catch(error => {
      console.log(error.response.status);
      setStatus(error.response.status);
      setOpen(true);
    });
  }

  return (
    <div>
      <nav className="navbar">
        <div className="container">
          <div className="navbar-brand">
            <h1 className="navbar-item is-size-4">Gating System</h1>
          </div>
          <div id="navbarMenu" className="navbar-menu">
            <div className="navbar-end is-align-items-center">
            </div>
          </div>
        </div>
      </nav>
      <section className="hero is-fullheight">
        <div className="faucet-hero-body">
          {boolMM ? 
            <div className="container has-text-centered main-content">
              <button
                className="button is-white connect-wallet"
                onClick={connectWallet}
              >
                <span className="is-link has-text-weight-bold">
                  {walletAddress && walletAddress.length > 0
                  ? `Wallet Connected Successfully `
                  : "Connect Wallet"}
                </span>
              </button>
              <div>
                <span className="is-link has-text-weight-bold">
                  {walletAddress && walletAddress.length > 0 && searchParams.get("sign")
                    ? 
                    <div>
                      Connected Address: {walletAddress.toString()} 
                        <div>
                          <button
                            className="button is-white connect-wallet"
                            onClick={signMessage}
                          >
                            <span className="is-link has-text-weight-bold">
                              Sign Message
                            </span>
                          </button>
                        </div>
                      </div>
                    : ""
                    }
                </span>
              </div>
            </div>
          : 
            <div>
              <span className="is-link has-text-weight-bold">
                Metamask not Installed
              </span>
            </div>
          }
        </div>
      </section>
      {status ?
        <Backdrop
          sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
          open={open}
          onClick={handleClose}
        >
          {status === 200 ? 
          <Alert severity="success">
            <AlertTitle>Success</AlertTitle>
            Status — <strong>200</strong>
          </Alert>
          :
          <Alert severity="error">
            <AlertTitle>Error</AlertTitle>
            
            <div>Status: <strong>{status}</strong> </div>
            <div>Contact an admin</div>
            
          </Alert>
          }
        </Backdrop>
      :""
      }
    </div>
  );
}

export default App;
