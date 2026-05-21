package com.takatime.jetbrains

import com.intellij.openapi.editor.Document
import com.intellij.openapi.editor.EditorFactory
import com.intellij.openapi.editor.event.DocumentEvent
import com.intellij.openapi.editor.event.DocumentListener
import com.intellij.openapi.editor.event.EditorFactoryEvent
import com.intellij.openapi.editor.event.EditorFactoryListener
import com.intellij.openapi.fileEditor.FileDocumentManagerListener
import java.util.WeakHashMap

// 1. Save Listener
class TakaTimeSaveListener : FileDocumentManagerListener {
    override fun beforeDocumentSaving(document: Document) {
        TakaTimeHeartbeat.handleHeartbeat(document)
    }
}

// 2. Typing Listener
class TakaTimeEditorFactoryListener : EditorFactoryListener {

    private val attachedDocuments = WeakHashMap<Document, Boolean>()

    private val typingListener = object : DocumentListener {
        override fun documentChanged(e: DocumentEvent) {
            TakaTimeHeartbeat.handleHeartbeat(e.document)
        }
    }

    override fun editorCreated(event: EditorFactoryEvent) {
        val document = event.editor.document

        TakaTimeBinaryManager.checkAndDownloadIfNeeded(
            event.editor.project
        )

        if (!attachedDocuments.containsKey(document)) {
            document.addDocumentListener(typingListener)
            attachedDocuments[document] = true
        }
    }

    override fun editorReleased(event: EditorFactoryEvent) {
        val document = event.editor.document

        val editors = EditorFactory
            .getInstance()
            .getEditors(document)

        if (editors.size <= 1 &&
            attachedDocuments.containsKey(document)
        ) {
            document.removeDocumentListener(typingListener)
            attachedDocuments.remove(document)
        }
    }
}