// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "protocol.proto" (syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message Diff
 */
export interface Diff {
    /**
     * @generated from protobuf oneof: root
     */
    root: {
        oneofKind: "document";
        /**
         * @generated from protobuf field: bool document = 1;
         */
        document: boolean;
    } | {
        oneofKind: "elementSelector";
        /**
         * @generated from protobuf field: string element_selector = 2;
         */
        elementSelector: string;
    } | {
        oneofKind: undefined;
    };
    /**
     * Position of each child
     *
     * @generated from protobuf field: repeated uint32 path_indicies = 3;
     */
    pathIndicies: number[];
    /**
     * @generated from protobuf field: DiffType diff_type = 4;
     */
    diffType: DiffType;
    /**
     * @generated from protobuf field: ContentType content_type = 5;
     */
    contentType: ContentType;
    /**
     * @generated from protobuf field: string contents = 6;
     */
    contents: string;
}
/**
 * @generated from protobuf message Diffs
 */
export interface Diffs {
    /**
     * @generated from protobuf field: repeated Diff values = 1;
     */
    values: Diff[];
}
/**
 * @generated from protobuf message SessionInfo
 */
export interface SessionInfo {
    /**
     * @generated from protobuf field: uint64 id = 1;
     */
    id: bigint;
}
/**
 * @generated from protobuf message ToClient
 */
export interface ToClient {
    /**
     * @generated from protobuf oneof: message
     */
    message: {
        oneofKind: "diffs";
        /**
         * @generated from protobuf field: Diffs diffs = 1;
         */
        diffs: Diffs;
    } | {
        oneofKind: "sessionInfo";
        /**
         * @generated from protobuf field: SessionInfo session_info = 2;
         */
        sessionInfo: SessionInfo;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message FromClient
 */
export interface FromClient {
    /**
     * @generated from protobuf field: FromClient.Type type = 1;
     */
    type: FromClient_Type;
    /**
     * @generated from protobuf field: repeated string ids = 2;
     */
    ids: string[];
    /**
     * @generated from protobuf field: bool selected = 3;
     */
    selected: boolean;
    /**
     * @generated from protobuf field: repeated string value_multi = 4;
     */
    valueMulti: string[];
    /**
     * @generated from protobuf field: map<string, string> data = 5;
     */
    data: {
        [key: string]: string;
    };
    /**
     * @generated from protobuf field: map<string, string> extra = 6;
     */
    extra: {
        [key: string]: string;
    };
    /**
     * @generated from protobuf field: FromClient.File file = 7;
     */
    file?: FromClient_File;
}
/**
 * @generated from protobuf message FromClient.File
 */
export interface FromClient_File {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: int64 SizeBytes = 2 [json_name = "SizeBytes"];
     */
    sizeBytes: bigint;
    /**
     * @generated from protobuf field: string mime_type = 3;
     */
    mimeType: string;
    /**
     * @generated from protobuf field: bytes data = 4;
     */
    data: Uint8Array;
    /**
     * @generated from protobuf field: uint32 total_files_index = 5;
     */
    totalFilesIndex: number;
    /**
     * @generated from protobuf field: uint32 total_file_count = 6;
     */
    totalFileCount: number;
}
/**
 * @generated from protobuf enum FromClient.Type
 */
export enum FromClient_Type {
    /**
     * @generated from protobuf enum value: TYPE_UNKNOWN = 0;
     */
    TYPE_UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: LOG = 1;
     */
    LOG = 1,
    /**
     * @generated from protobuf enum value: EVENT = 2;
     */
    EVENT = 2
}
// import "vtproto/ext.proto";

/**
 * @generated from protobuf enum DiffType
 */
export enum DiffType {
    /**
     * @generated from protobuf enum value: DIFF_TYPE_UNKNOWN = 0;
     */
    DIFF_TYPE_UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: CREATE = 1;
     */
    CREATE = 1,
    /**
     * @generated from protobuf enum value: UPDATE = 2;
     */
    UPDATE = 2,
    /**
     * @generated from protobuf enum value: DELETE = 3;
     */
    DELETE = 3
}
/**
 * @generated from protobuf enum ContentType
 */
export enum ContentType {
    /**
     * @generated from protobuf enum value: CONTENT_TYPE_UNKNOWN = 0;
     */
    CONTENT_TYPE_UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: TEXT = 1;
     */
    TEXT = 1,
    /**
     * @generated from protobuf enum value: HTML = 2;
     */
    HTML = 2,
    /**
     * @generated from protobuf enum value: ATTRIBUTE = 3;
     */
    ATTRIBUTE = 3,
    /**
     * @generated from protobuf enum value: EMPTY_FROM_REMOVAL = 4;
     */
    EMPTY_FROM_REMOVAL = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class Diff$Type extends MessageType<Diff> {
    constructor() {
        super("Diff", [
            { no: 1, name: "document", kind: "scalar", oneof: "root", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "element_selector", kind: "scalar", oneof: "root", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "path_indicies", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 13 /*ScalarType.UINT32*/ },
            { no: 4, name: "diff_type", kind: "enum", T: () => ["DiffType", DiffType] },
            { no: 5, name: "content_type", kind: "enum", T: () => ["ContentType", ContentType] },
            { no: 6, name: "contents", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Diff>): Diff {
        const message = { root: { oneofKind: undefined }, pathIndicies: [], diffType: 0, contentType: 0, contents: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Diff>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Diff): Diff {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool document */ 1:
                    message.root = {
                        oneofKind: "document",
                        document: reader.bool()
                    };
                    break;
                case /* string element_selector */ 2:
                    message.root = {
                        oneofKind: "elementSelector",
                        elementSelector: reader.string()
                    };
                    break;
                case /* repeated uint32 path_indicies */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.pathIndicies.push(reader.uint32());
                    else
                        message.pathIndicies.push(reader.uint32());
                    break;
                case /* DiffType diff_type */ 4:
                    message.diffType = reader.int32();
                    break;
                case /* ContentType content_type */ 5:
                    message.contentType = reader.int32();
                    break;
                case /* string contents */ 6:
                    message.contents = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Diff, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool document = 1; */
        if (message.root.oneofKind === "document")
            writer.tag(1, WireType.Varint).bool(message.root.document);
        /* string element_selector = 2; */
        if (message.root.oneofKind === "elementSelector")
            writer.tag(2, WireType.LengthDelimited).string(message.root.elementSelector);
        /* repeated uint32 path_indicies = 3; */
        if (message.pathIndicies.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.pathIndicies.length; i++)
                writer.uint32(message.pathIndicies[i]);
            writer.join();
        }
        /* DiffType diff_type = 4; */
        if (message.diffType !== 0)
            writer.tag(4, WireType.Varint).int32(message.diffType);
        /* ContentType content_type = 5; */
        if (message.contentType !== 0)
            writer.tag(5, WireType.Varint).int32(message.contentType);
        /* string contents = 6; */
        if (message.contents !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.contents);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message Diff
 */
export const Diff = new Diff$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Diffs$Type extends MessageType<Diffs> {
    constructor() {
        super("Diffs", [
            { no: 1, name: "values", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Diff }
        ]);
    }
    create(value?: PartialMessage<Diffs>): Diffs {
        const message = { values: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Diffs>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Diffs): Diffs {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated Diff values */ 1:
                    message.values.push(Diff.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Diffs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated Diff values = 1; */
        for (let i = 0; i < message.values.length; i++)
            Diff.internalBinaryWrite(message.values[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message Diffs
 */
export const Diffs = new Diffs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SessionInfo$Type extends MessageType<SessionInfo> {
    constructor() {
        super("SessionInfo", [
            { no: 1, name: "id", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<SessionInfo>): SessionInfo {
        const message = { id: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SessionInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SessionInfo): SessionInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 id */ 1:
                    message.id = reader.uint64().toBigInt();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SessionInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 id = 1; */
        if (message.id !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message SessionInfo
 */
export const SessionInfo = new SessionInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ToClient$Type extends MessageType<ToClient> {
    constructor() {
        super("ToClient", [
            { no: 1, name: "diffs", kind: "message", oneof: "message", T: () => Diffs },
            { no: 2, name: "session_info", kind: "message", oneof: "message", T: () => SessionInfo }
        ]);
    }
    create(value?: PartialMessage<ToClient>): ToClient {
        const message = { message: { oneofKind: undefined } };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ToClient>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ToClient): ToClient {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* Diffs diffs */ 1:
                    message.message = {
                        oneofKind: "diffs",
                        diffs: Diffs.internalBinaryRead(reader, reader.uint32(), options, (message.message as any).diffs)
                    };
                    break;
                case /* SessionInfo session_info */ 2:
                    message.message = {
                        oneofKind: "sessionInfo",
                        sessionInfo: SessionInfo.internalBinaryRead(reader, reader.uint32(), options, (message.message as any).sessionInfo)
                    };
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ToClient, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* Diffs diffs = 1; */
        if (message.message.oneofKind === "diffs")
            Diffs.internalBinaryWrite(message.message.diffs, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* SessionInfo session_info = 2; */
        if (message.message.oneofKind === "sessionInfo")
            SessionInfo.internalBinaryWrite(message.message.sessionInfo, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message ToClient
 */
export const ToClient = new ToClient$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FromClient$Type extends MessageType<FromClient> {
    constructor() {
        super("FromClient", [
            { no: 1, name: "type", kind: "enum", T: () => ["FromClient.Type", FromClient_Type] },
            { no: 2, name: "ids", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "selected", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "value_multi", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "data", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 6, name: "extra", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 7, name: "file", kind: "message", T: () => FromClient_File }
        ]);
    }
    create(value?: PartialMessage<FromClient>): FromClient {
        const message = { type: 0, ids: [], selected: false, valueMulti: [], data: {}, extra: {} };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<FromClient>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: FromClient): FromClient {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* FromClient.Type type */ 1:
                    message.type = reader.int32();
                    break;
                case /* repeated string ids */ 2:
                    message.ids.push(reader.string());
                    break;
                case /* bool selected */ 3:
                    message.selected = reader.bool();
                    break;
                case /* repeated string value_multi */ 4:
                    message.valueMulti.push(reader.string());
                    break;
                case /* map<string, string> data */ 5:
                    this.binaryReadMap5(message.data, reader, options);
                    break;
                case /* map<string, string> extra */ 6:
                    this.binaryReadMap6(message.extra, reader, options);
                    break;
                case /* FromClient.File file */ 7:
                    message.file = FromClient_File.internalBinaryRead(reader, reader.uint32(), options, message.file);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    private binaryReadMap5(map: FromClient["data"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof FromClient["data"] | undefined, val: FromClient["data"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = reader.string();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field FromClient.data");
            }
        }
        map[key ?? ""] = val ?? "";
    }
    private binaryReadMap6(map: FromClient["extra"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof FromClient["extra"] | undefined, val: FromClient["extra"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = reader.string();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field FromClient.extra");
            }
        }
        map[key ?? ""] = val ?? "";
    }
    internalBinaryWrite(message: FromClient, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* FromClient.Type type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* repeated string ids = 2; */
        for (let i = 0; i < message.ids.length; i++)
            writer.tag(2, WireType.LengthDelimited).string(message.ids[i]);
        /* bool selected = 3; */
        if (message.selected !== false)
            writer.tag(3, WireType.Varint).bool(message.selected);
        /* repeated string value_multi = 4; */
        for (let i = 0; i < message.valueMulti.length; i++)
            writer.tag(4, WireType.LengthDelimited).string(message.valueMulti[i]);
        /* map<string, string> data = 5; */
        for (let k of Object.keys(message.data))
            writer.tag(5, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k).tag(2, WireType.LengthDelimited).string(message.data[k]).join();
        /* map<string, string> extra = 6; */
        for (let k of Object.keys(message.extra))
            writer.tag(6, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k).tag(2, WireType.LengthDelimited).string(message.extra[k]).join();
        /* FromClient.File file = 7; */
        if (message.file)
            FromClient_File.internalBinaryWrite(message.file, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message FromClient
 */
export const FromClient = new FromClient$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FromClient_File$Type extends MessageType<FromClient_File> {
    constructor() {
        super("FromClient.File", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "SizeBytes", kind: "scalar", jsonName: "SizeBytes", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "mime_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 5, name: "total_files_index", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 6, name: "total_file_count", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
    create(value?: PartialMessage<FromClient_File>): FromClient_File {
        const message = { name: "", sizeBytes: 0n, mimeType: "", data: new Uint8Array(0), totalFilesIndex: 0, totalFileCount: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<FromClient_File>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: FromClient_File): FromClient_File {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* int64 SizeBytes = 2 [json_name = "SizeBytes"];*/ 2:
                    message.sizeBytes = reader.int64().toBigInt();
                    break;
                case /* string mime_type */ 3:
                    message.mimeType = reader.string();
                    break;
                case /* bytes data */ 4:
                    message.data = reader.bytes();
                    break;
                case /* uint32 total_files_index */ 5:
                    message.totalFilesIndex = reader.uint32();
                    break;
                case /* uint32 total_file_count */ 6:
                    message.totalFileCount = reader.uint32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: FromClient_File, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* int64 SizeBytes = 2 [json_name = "SizeBytes"]; */
        if (message.sizeBytes !== 0n)
            writer.tag(2, WireType.Varint).int64(message.sizeBytes);
        /* string mime_type = 3; */
        if (message.mimeType !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.mimeType);
        /* bytes data = 4; */
        if (message.data.length)
            writer.tag(4, WireType.LengthDelimited).bytes(message.data);
        /* uint32 total_files_index = 5; */
        if (message.totalFilesIndex !== 0)
            writer.tag(5, WireType.Varint).uint32(message.totalFilesIndex);
        /* uint32 total_file_count = 6; */
        if (message.totalFileCount !== 0)
            writer.tag(6, WireType.Varint).uint32(message.totalFileCount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message FromClient.File
 */
export const FromClient_File = new FromClient_File$Type();
