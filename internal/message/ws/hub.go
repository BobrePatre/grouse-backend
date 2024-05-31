package ws

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"sync"
)

type Client struct {
	TelegramId string
	Conn       *websocket.Conn
}

type Hub struct {
	upgrader *websocket.Upgrader
	logger   *slog.Logger
	clients  map[string]Client
	mx       sync.Mutex
}

func NewHub(upgrader *websocket.Upgrader, logger *slog.Logger) *Hub {
	return &Hub{
		upgrader: upgrader,
		logger:   logger,
		clients:  make(map[string]Client),
	}
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error("Failed to upgrade websocket connection", "error", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			h.logger.Error("Failed to close websocket connection", "error", err)
		}
	}(conn)
	telegramId := r.URL.Query().Get("telegramId")
	if telegramId == "" {
		h.logger.Error("Failed to get telegram id from request")
		return
	}
	h.mx.Lock()
	h.clients[telegramId] = Client{
		TelegramId: telegramId,
		Conn:       conn,
	}
	h.mx.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.logger.Error("Failed to read message", "error", err)
			delete(h.clients, telegramId)
			break
		}
		h.logger.Info("Received message", "telegramId", telegramId, "msg", msg)
	}
}

func RegisterWsHub(srv *echo.Echo, hub *Hub) error {
	srv.Any("/ws", echo.WrapHandler(hub))
	return nil
}

func (h *Hub) GetClient(telegramId string) (Client, error) {
	h.mx.Lock()
	defer h.mx.Unlock()
	client, ok := h.clients[telegramId]
	if !ok {
		return client, errors.New("client not found")
	}
	return client, nil
}
