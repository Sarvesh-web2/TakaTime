package com.takatime.jetbrains

import com.intellij.openapi.editor.Document
import com.intellij.openapi.editor.EditorFactory
import com.intellij.openapi.editor.event.DocumentEvent
import com.intellij.openapi.editor.event.DocumentListener
import com.intellij.openapi.editor.event.EditorFactoryEvent
import com.intellij.openapi.editor.event.EditorFactoryListener
import com.intellij.openapi.fileEditor.FileDocumentManager
import com.intellij.openapi.fileEditor.FileDocumentManagerListener

// 1. The Save Listener (Ctrl+S)
class TakaTimeSaveListener : FileDocumentManagerListener {
    override fun beforeDocumentSaving(document: Document) {
        TakaTimeHeartbeat.handleHeartbeat(document)
    }
}

// 2. The Typing & Boot-up Listener
class TakaTimeEditorFactoryListener : EditorFactoryListener {
    private val typingListener = object : DocumentListener {
        override fun documentChanged(e: DocumentEvent) {
            TakaTimeHeartbeat.handleHeartbeat(e.document)
        }
    }

    override fun editorCreated(event: EditorFactoryEvent) {
        val document = event.editor.document
        val editors = EditorFactory.getInstance().getEditors(document)

        // Asynchronously check if binaries need downloading when an editor opens!
        TakaTimeBinaryManager.checkAndDownloadIfNeeded(event.editor.project)

        if (editors.size == 1) {
            document.addDocumentListener(typingListener)
        }
    }

    override fun editorReleased(event: EditorFactoryEvent) {
        val document = event.editor.document
        val editors = EditorFactory.getInstance().getEditors(document)

        if (editors.size <= 1) {
            document.removeDocumentListener(typingListener)
        }
    }
}