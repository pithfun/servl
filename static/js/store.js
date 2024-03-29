/**
 * Constants
 */
const APP_NAME = '__tiny_blob__';
const __provider__ = window.phantom?.solana;

/**
 * Alpine store
 */
document.addEventListener('alpine:init', () => {
  Alpine.store(APP_NAME, {
    ...__provider__,
    app_name: APP_NAME,
    pub_key: null,
    is_phantom_installed: false,
    is_connected: false,
    init_wallet(pub_key) {
      this.is_phantom_installed = true;
      this.is_connected = true;
      this.pub_key = pub_key;
    },
    clean_up() {
      this.is_connected = false;
      this.pub_key = null;
    },
    init() {
      if (__provider__?.isPhantom) {
        __provider__.on('connect', (pub_key) => {
          if (pub_key) {
            this.init_wallet();
          } else {
            this.clean_up();
          }
        });

        __provider__.on('disconnect', () => {
          this.clean_up();
        });

        __provider__.on('accountChanged', (pub_key) => {
          if (pub_key) {
            let pub_key_str = pub_key.toBase58();
            this.init_wallet(pub_key_str);
            window.toast('Switched account', {
              type: 'info',
              description: `Switched to ${pub_key_str.substring(0, 5)}…${pub_key_str.substring(pub_key_str.length - 5)}`,
              position: 'bottom-left',
            });
          } else {
            // Account is not connected.
            // Prompt user to connect.
            this.clean_up();
            this.connect_wallet();
          }
        });

        // Attempt to eagerly connect
        __provider__.connect({ onlyIfTrusted: true }).catch((e) => {
          /* ignore error */
        });
      } else {
        console.warn('Phantom wallet is not installed');
      }
    },
    async connect_wallet() {
      if (!this.is_connected) {
        try {
          const res = await this.connect();
          this.init_wallet(res.publicKey.toBase58());

          window.toast('Connected', {
            type: 'success',
            description: 'Wallet has been connected.',
            position: 'bottom-left',
          });
        } catch (e) {
          let msg = '';
          switch (true) {
            case e instanceof Error && e.message.includes('rejected'):
              msg = 'Wallet connection cancelled.';
              break;
            default:
              msg = 'Could not connect to your wallet.';
              break;
          }
          window.toast('Unsuccessful', {
            type: 'warning',
            description: msg,
            position: 'bottom-left',
          });
        }
      }
    },
    async disconnect_wallet() {
      if (this.is_connected && this.is_phantom_installed) {
        try {
          this.clean_up();
          window.toast('Disconnected', {
            type: 'info',
            description: 'Wallet has been disconnected.',
            position: 'bottom-left',
          });
        } catch (e) {
          console.error(e);
        }
      }
    },
  });
});
