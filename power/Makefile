CFLAGS = `pkg-config --cflags --libs glib-2.0 gtk+-3.0 gnome-desktop-3.0 gudev-1.0`

gsd-power-manager:gsd-power-manager.c gpm-common.c gsd-backlight-helper.c gsd-backlight-linux.c gsd-power-plugin.c
	$(CC) -o $@ $^ $(CFLAGS)
