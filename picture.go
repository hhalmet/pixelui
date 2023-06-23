package pixelui

import (
	"sync/atomic"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go/v4"
)

func (ui *UI) AddOrSetPicture(name string, pic pixel.Picture) imgui.TextureID {
	id := ui.nextPictureId()
	drawable := &renderSource{
		picture:    pic,
		shaderTris: pixelgl.NewGLTriangles(ui.shader, pixel.MakeTrianglesData(0)),
	}
	ui.pictureDrawables[name] = drawable
	ui.pictureNames[id] = name

	return imgui.TextureID(id + PICTURE_TEXTUREID_OFFSET)
}

func (ui *UI) RemovePicture(id imgui.TextureID) {
	ui.pictureDrawables[ui.pictureNames[int(id)]] = nil
}

func (ui *UI) nextPictureId() int {
	return int(atomic.AddInt32(&ui.lastPictureId, 1))
}
